# Lab 7 of GUI Programming which covers the UI-Building-Pattern

Task was to implement an event registration feature following the form-v1 template we had from a previous class.

Some topics covered were:
- Form Validation
	- Event Date is:
		- `HTML type="date"`
		- Must be a future date
	- Number of Tickets is:
		- `HTML type="number"`
		- Must be b/w 1 and 5
	- Terms and conditions is:
		- `HTML type="checkbox"`
		- Must be checked to proceed
- Extending APIs
	- to accept new payload
- Observer Pattern
	- subscribing it the form input event and emitting notifications

---
The primary changes made from form-v1 template project was:
- `main.go`
	- Struct expansion to include EventDate and Tickets
	- Time validation using `time` package and using time.Parse and Go reference layout `2006-01-02`
	- Validations for new data
- `render.js`
	- `minDate` variable inside renderForm to set minimum allowed date to after today on initial render and serve
	- Data Transformation:
		- using `parseInt` for tickets to treat the string as an int for validation
- `handleSubmit`
	- added new data into payload
		- including constraint date and formatting for dates
- `renderForm()`:
	- setting min date for the calendar dropdown
	- including new form UI generation for event Registration data

---
Some concepts that clicked:
- Extending the API to accept more or different types of form data
- implementing validation on the backend

Some challenges faced:
- Dynamic UI building
	- More so with implementing better looking invalid form notification, I just settled on hovers but I should have looked more on in UI alerts and logging using GO's or JS's standard libraries.
- working with GO
	- much more strict than the vanilla JavaScript we've been used to thus far, but also more robust in error checking as it doesn't run if backend logic is faulty

For future projects:
- Form validation
- working more with GO as a backend language to drive JS applications
- working more with and on APIs
- Dynamic UI Building pattern
