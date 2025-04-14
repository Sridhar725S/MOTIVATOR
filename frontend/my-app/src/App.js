import { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [quote, setQuote] = useState("Loading wisdom...");

  const getQuote = async () => {
    const res = await fetch('https://motivator-73sm.onrender.com/api/quote');
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
