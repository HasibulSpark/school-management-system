const studentForm = document.getElementById("student-form");
const formMessage = document.getElementById("form-message");
const studentsTableBody = document.getElementById("students-table-body");

function showMessage(message, isError = false) {
  formMessage.textContent = message;
  formMessage.className = isError ? "message error-message" : "message success-message";
}

function createCell(row, value) {
  const cell = document.createElement("td");
  cell.textContent = value || "-";
  row.appendChild(cell);
}

async function loadStudents() {
  try {
    const response = await fetch("/api/students");

    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}`);
    }

    const students = await response.json();
    studentsTableBody.replaceChildren();

    students.forEach((student) => {
      const row = document.createElement("tr");

      createCell(row, student.id);
      createCell(row, student.name);
      createCell(row, student.class);
      createCell(row, student.section);
      createCell(row, student.roll);
      createCell(row, student.guardianName);
      createCell(row, student.guardianPhone);

      studentsTableBody.appendChild(row);
    });
  } catch (error) {
    console.error("Could not load students:", error);
    studentsTableBody.innerHTML =
      '<tr><td colspan="7">Could not load students. Make sure the Go API is running.</td></tr>';
  }
}

studentForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  showMessage("Registering student...");

  const student = {
    name: document.getElementById("name").value,
    dateOfBirth: document.getElementById("date-of-birth").value,
    class: document.getElementById("class").value,
    section: document.getElementById("section").value,
    roll: document.getElementById("roll").value,
    guardianName: document.getElementById("guardian-name").value,
    guardianPhone: document.getElementById("guardian-phone").value,
    address: document.getElementById("address").value,
  };

  try {
    const response = await fetch("/api/students", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(student),
    });

    if (!response.ok) {
      const errorMessage = await response.text();
      throw new Error(errorMessage);
    }

    studentForm.reset();
    showMessage("Student registered successfully.");
    await loadStudents();
  } catch (error) {
    console.error("Could not register student:", error);
    showMessage(error.message || "Could not register student.", true);
  }
});

loadStudents();