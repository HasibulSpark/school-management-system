const teacherForm = document.getElementById("teacher-form");
const teacherFormMessage = document.getElementById("form-message");
const teachersTableBody = document.getElementById("teachers-table-body");

function showTeacherMessage(message, isError = false) {
  teacherFormMessage.textContent = message;
  teacherFormMessage.className = isError
    ? "message error-message"
    : "message success-message";
}

function addTeacherCell(row, value) {
  const cell = document.createElement("td");
  cell.textContent = value || "-";
  row.appendChild(cell);
}

function formatSalary(salary) {
  const amount = Number(salary || 0);

  if (amount === 0) {
    return "-";
  }

  return `BDT ${amount.toLocaleString()}`;
}

async function loadTeachers() {
  try {
    const response = await fetch("/api/teachers");

    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}`);
    }

    const teachers = await response.json();
    teachersTableBody.replaceChildren();

    teachers.forEach((teacher) => {
      const row = document.createElement("tr");

      addTeacherCell(row, teacher.id);
      addTeacherCell(row, teacher.name);
      addTeacherCell(row, teacher.employeeId);
      addTeacherCell(row, teacher.subject);
      addTeacherCell(row, teacher.assignedClass);
      addTeacherCell(row, teacher.phone);
      addTeacherCell(row, teacher.email);
      addTeacherCell(row, formatSalary(teacher.monthlySalary));

      teachersTableBody.appendChild(row);
    });
  } catch (error) {
    console.error("Could not load teachers:", error);
    teachersTableBody.innerHTML =
      '<tr><td colspan="8">Could not load teachers. Make sure the Go API is running.</td></tr>';
  }
}

teacherForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  showTeacherMessage("Registering teacher...");

  const teacher = {
    name: document.getElementById("name").value,
    employeeId: document.getElementById("employee-id").value,
    phone: document.getElementById("phone").value,
    email: document.getElementById("email").value,
    subject: document.getElementById("subject").value,
    assignedClass: document.getElementById("assigned-class").value,
    joiningDate: document.getElementById("joining-date").value,
    monthlySalary: Number(document.getElementById("monthly-salary").value || 0),
  };

  try {
    const response = await fetch("/api/teachers", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(teacher),
    });

    if (!response.ok) {
      const errorMessage = await response.text();
      throw new Error(errorMessage);
    }

    teacherForm.reset();
    showTeacherMessage("Teacher registered successfully.");
    await loadTeachers();
  } catch (error) {
    console.error("Could not register teacher:", error);
    showTeacherMessage(error.message || "Could not register teacher.", true);
  }
});

loadTeachers();