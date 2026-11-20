const attendanceForm = document.getElementById("attendance-form");
const attendanceFormMessage = document.getElementById("form-message");
const attendanceTableBody = document.getElementById("attendance-table-body");

function showAttendanceMessage(message, isError = false) {
  attendanceFormMessage.textContent = message;
  attendanceFormMessage.className = isError
    ? "message error-message"
    : "message success-message";
}

function addAttendanceCell(row, value) {
  const cell = document.createElement("td");
  cell.textContent = value || "-";
  row.appendChild(cell);
}

function createStatusBadge(status) {
  const cell = document.createElement("td");
  const badge = document.createElement("span");

  badge.textContent = status;
  badge.className = `attendance-status ${status}`;

  cell.appendChild(badge);
  return cell;
}

async function loadAttendance() {
  try {
    const response = await fetch("/api/attendance");

    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}`);
    }

    const records = await response.json();
    attendanceTableBody.replaceChildren();

    records.forEach((record) => {
      const row = document.createElement("tr");

      addAttendanceCell(row, record.id);
      addAttendanceCell(row, record.date);
      addAttendanceCell(row, record.studentName);
      addAttendanceCell(row, record.class);
      addAttendanceCell(row, record.section);
      addAttendanceCell(row, record.roll);
      row.appendChild(createStatusBadge(record.status));

      attendanceTableBody.appendChild(row);
    });
  } catch (error) {
    console.error("Could not load attendance:", error);
    attendanceTableBody.innerHTML =
      '<tr><td colspan="7">Could not load attendance. Make sure the Go API is running.</td></tr>';
  }
}

attendanceForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  showAttendanceMessage("Saving attendance...");

  const record = {
    date: document.getElementById("date").value,
    studentName: document.getElementById("student-name").value,
    class: document.getElementById("class").value,
    section: document.getElementById("section").value,
    roll: document.getElementById("roll").value,
    status: document.getElementById("status").value,
  };

  try {
    const response = await fetch("/api/attendance", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(record),
    });

    if (!response.ok) {
      const errorMessage = await response.text();
      throw new Error(errorMessage);
    }

    attendanceForm.reset();
    showAttendanceMessage("Attendance saved successfully.");
    await loadAttendance();
  } catch (error) {
    console.error("Could not save attendance:", error);
    showAttendanceMessage(error.message || "Could not save attendance.", true);
  }
});

document.getElementById("date").value = new Date().toISOString().slice(0, 10);

loadAttendance();