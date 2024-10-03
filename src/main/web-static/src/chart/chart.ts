import { Chart, registerables } from "chart.js";

Chart.register(...registerables);

const chartCanvas = document.getElementById("testChart") as HTMLCanvasElement;

new Chart(chartCanvas, {
    type: "pie",
    data: { datasets: [{ data: [10, 20, 30], label: "Test dataset" }] },
    options: { responsive: true, maintainAspectRatio: true },  
});
