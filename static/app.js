document.addEventListener("DOMContentLoaded", () => {
    const today = new Date();
    const formattedDate = today.toISOString().split('T')[0]; // YYYY-MM-DD
    const formattedMonth = today.toISOString().slice(0, 7); // YYYY-MM
  
    const dateInput = document.getElementById("date");
    const portfolioDateInput = document.getElementById("portfolio-date");
  
    if (dateInput) dateInput.value = formattedDate;
    if (portfolioDateInput) portfolioDateInput.value = formattedMonth;
  });
  
  // Handle UI switching
  const registerSpendingBtn = document.getElementById("register-spending-btn");
  const portfolioViewBtn = document.getElementById("portfolio-view-btn");
  const spendingForm = document.getElementById("spending-form");
  const portfolioForm = document.getElementById("portfolio-form");
  
  registerSpendingBtn.addEventListener("click", () => {
    spendingForm.style.display = "block";
    portfolioForm.style.display = "none";
  });
  
  portfolioViewBtn.addEventListener("click", () => {
    spendingForm.style.display = "none";
    portfolioForm.style.display = "block";
  });
  
  // Handle Spending Form Submission
  document.getElementById("spending-form-element").addEventListener("submit", async (event) => {
    event.preventDefault();
    const formData = new FormData(event.target);
  
    const payload = {
      date: formData.get("date"),
      name: formData.get("name"),
      category: formData.get("category"),
      amount: parseInt(formData.get("amount")),
      note: formData.get("note")
    };
  
    try {
      const response = await fetch("${CONFIG.API_BASE_URL}/api/transaction", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
      });
  
      if (response.ok) {
        alert("Transaction saved successfully!");
      } else {
        const errorMsg = await response.text();
        alert(`Failed to save transaction: ${errorMsg}`);
      }
    } catch (error) {
      alert(`Error occurred while saving transaction: ${error.message}`);
    }
  });
  
  // Handle Portfolio Form Submission
  document.getElementById("portfolio-form-element").addEventListener("submit", async (event) => {
    event.preventDefault();
    const date = document.getElementById("portfolio-date").value;
  
    try {
      const response = await fetch(`${CONFIG.API_BASE_URL}/api/portfolio?date=${date}`, {
        method: "GET",
        headers: {
          "Accept": "text/plain"
        }
      });
  
      if (response.ok) {
        const result = await response.text();
        document.getElementById("portfolio-result").textContent = result;
      } else {
        const errorMsg = await response.text();
        alert(`Failed to fetch portfolio: ${errorMsg}`);
      }
    } catch (error) {
      alert(`Error occurred while fetching portfolio: ${error.message}`);
    }
  });
  