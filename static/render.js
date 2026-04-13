import { state } from "./state.js";
import { emitter } from "./modules/event-emitter.js";

function handleSubmit(e) {
  e.preventDefault();

  const data = new FormData(e.target);
  const dateString = data.get("event_date");
  const eventDate = new Date(dateString);
  // Validation: Compare event date with today's date
  const today = new Date();
  today.setHours(0, 0, 0, 0); // Reset time for accurate date comparison

  if (eventDate <= today) {
    alert("Event must be in the future.");
    return;
  }
  // update payload object
  const payload = {
    name: data.get("name").trim(),
    email: data.get("email").trim(),
    event_date: dateString,
    tickets: parseInt(data.get("tickets")),
    agreed: data.get("terms") === "on",
  };

  emitter.emit("users:submit", payload);
}

function renderForm() {
  // For render form, restructure to include event registration form while keeping old user form fields
  // Constraint must be added to make the minimum date the current date from today const
  const minDate = new Date().toISOString().split("T")[0];
  return `
    <div class="form-card">
      <h2>Event Registration</h2>
      <form id="user-form" class="ui-form">

        <input
          class="ui-input"
          name="name"
          placeholder="Full name"
          required
        />

        <input
          class="ui-input"
          name="email"
          type="email"
          placeholder="Email address"
          required
        />

      <label> Event Date (Future Only)</label>
      <input
        class="ui-input"
        name="event_date"
        type="date"
        min="${minDate}"
        required
      />

      <label>Tickets (1-5)</label>
      <input
        class="ui-input"
        name="tickets"
        type="number"
        min="1"
        max="5"
        value="1"
        required
      />

      <div style="display:flex; gap:10px; align-items:center; margin: 10px 0;">
        <input type="checkbox" name="terms" id="terms" required>
        <label for="terms" style="font-size:12px">I agree to Terms & Conditions</label>
      </div>

      <button class="ui-btn" type="submit">Submit</button>
    </form>
  </div>`;
}

// added div before submit is for Checkbox
export function render() {
  const app = document.querySelector("#app");
  if (app) {
    app.innerHTML = renderForm();
  }

  const form = document.querySelector("#user-form");
  if (form) {
    form.addEventListener("submit", handleSubmit);
  }
}
