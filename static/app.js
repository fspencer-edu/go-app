const sendBtn = document.getElementById("sendBtn");
const nameInput = document.getElementById("nameInput");
const result = document.getElementById("result");

sendBtn.addEventListener("click", async () => {
  const name = nameInput.value.trim();

  try {
    const response = await fetch("/api/greet", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    });

    if (!response.ok) {
      throw new Error("Request failed");
    }

    const data = await response.json();
    result.textContent = data.message;
  } catch (error) {
    result.textContent = "Something went wrong.";
    console.error(error);
  }
});