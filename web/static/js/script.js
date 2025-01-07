/*
 * TODO сделать нормально
 */
var gridSizeCols = 0;
var gridSizeRows = 0;

const gridElement = document.getElementById('grid');

// Загрузка сетки с сервера
function loadGrid() {
  fetch('/api/v1/state')
    .then(response => response.json())
    .then(data => renderGrid(data));
}

/*
 * TODO сделать нормально
 */
function gridSize() {
  fetch('/api/v1/size')
  .then(response => response.json())
  .then(data => setGridSize(data));
}

/*
 * TODO сделать нормально
 */
function setGridSize(data) {
  data.forEach((val, i) => {
    if (i == 0) {
      gridSizeCols = val
    } else {
      gridSizeRows = val
    }
  })
}


// Отрисовка сетки
function renderGrid(grid) {
  gridElement.innerHTML = '';
  gridElement.style.gridTemplateColumns = `repeat(${gridSizeCols}, 20px)`;
  gridElement.style.gridTemplateRows = `repeat(${gridSizeRows}, 20px)`;

  grid.forEach((row, i) => {
    row.forEach((cell, j) => {
      const cellElement = document.createElement('div');
      cellElement.classList.add('cell');
      if (cell === 1) {
        cellElement.classList.add('alive');
      }
      cellElement.addEventListener('click', () => toggleCell(i, j));
      gridElement.appendChild(cellElement);
    });
  });
}

// Изменение состояния клетки
function toggleCell(row, col) {
  fetch(`/api/v1/toggle?row=${row}&col=${col}`)
    .then(response => response.json())
    .then(data => renderGrid(data));
}

// Переход к следующему поколению
function nextGeneration() {
  fetch('/api/v1/next')
    .then(response => response.json())
    .then(data => renderGrid(data)); 
}

// Рандомное состояние
function seedState() {
  fetch('/api/v1/seed?fill=30')
    .then(response => response.json())
    .then(data => renderGrid(data)); 
}

// Первоначальная загрузка
gridSize();
loadGrid();