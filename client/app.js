document.addEventListener('DOMContentLoaded', function() {
  const num = document.getElementById('num');
  const advice = document.getElementById('advice');
  const dice = document.getElementById('dice');

  const URL = "http://localhost:8080";

  dice.addEventListener('click', function() {
      fetchNewQuote();
  });

  async function fetchNewQuote() {
      try {
          const response = await fetch(`${URL}/api/quote`);
          const data = await response.json();

          num.textContent = '#' + data.id;
          advice.textContent = data.content;
      } catch (err) {
          console.error('Error fetching quote:', err);
      }
  }
});
