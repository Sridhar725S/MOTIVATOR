import { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [quote, setQuote] = useState("Loading wisdom...");

  const getQuote = async () => {
    const res = await fetch('http://localhost:8080/api/quote');
    const data = await res.json();
    setQuote(data.quote);
  };

  useEffect(() => {
    getQuote();
  }, []);

  return (
    <div className="App" style={{ textAlign: 'center', marginTop: '2rem' }}>
      <h1>Motivator 3000 💬</h1>
      <p>{quote}</p>
      <button onClick={getQuote}>Give me motivation 🔁</button>
    </div>
  );
}

export default App;
