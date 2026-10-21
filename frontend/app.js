async function loadDashboardSummary() {
  try {
    const response = await fetch("/api/dashboard/summary");

    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}`);
    }

    const summary = await response.json();

    document.getElementById("total-students").textContent = summary.totals.students;
    document.getElementById("total-teachers").textContent = summary.totals.teachers;
    document.getElementById("total-staff").textContent = summary.totals.staff;
    document.getElementById("total-classes").textContent = summary.totals.classes;
    document.getElementById("total-sections").textContent = summary.totals.sections;

    document.getElementById("attendance-present").textContent = summary.attendance.present;
    document.getElementById("attendance-absent").textContent = summary.attendance.absent;
    document.getElementById("attendance-late").textContent = summary.attendance.late;

    document.getElementById("unpaid-students").textContent =
      summary.unpaidFees.students;
    document.getElementById("unpaid-amount").textContent =
      `${summary.unpaidFees.currency} ${summary.unpaidFees.amount}`;

    const activitiesList = document.getElementById("recent-activities");
    activitiesList.replaceChildren();

    summary.recentActivities.forEach((activity) => {
      const item = document.createElement("li");
      item.textContent = `${activity.time} — ${activity.description}`;
      activitiesList.appendChild(item);
    });
  } catch (error) {
    console.error("Could not load dashboard summary:", error);
    document.getElementById("recent-activities").textContent =
      "Could not load dashboard data. Make sure the Go API is running.";
  }
}

loadDashboardSummary();