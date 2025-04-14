import { test, expect } from '@playwright/test';

test('Motivator app loads and changes quote', async ({ page }) => {
  await page.goto('http://localhost:3000');
  
  const quote = await page.locator('p').textContent();
  await page.click('button');
  
  const newQuote = await page.locator('p').textContent();
  expect(newQuote).not.toEqual(quote);
});
