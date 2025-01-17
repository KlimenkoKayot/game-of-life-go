/*
 * TODO сделать нормально
 */
var gridSizeCols = 0;
var gridSizeRows = 0;
var autoGenerating = false;
var addNumsOnState = false;

const gridElement = document.getElementById('grid');
var numNeighbours;

// спим
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

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
      gridSizeRows = val
    } else {
      gridSizeCols = val
    }
  })
}

function realRenderGrid(grid) {
  gridElement.innerHTML = '';
  gridElement.style.gridTemplateColumns = `repeat(${gridSizeCols}, 20px)`;
  gridElement.style.gridTemplateRows = `repeat(${gridSizeRows}, 20px)`;
  grid.forEach((row, i) => {
    row.forEach((cell, j) => {
      const cellElement = document.createElement('div');
      cellElement.classList.add('cell');
      cellElement.id = `${i}_${j}`;
      if (cell === 1) {
        cellElement.classList.add('alive');
      }
      cellElement.addEventListener('click', () => toggleCell(i, j));
      if (addNumsOnState) {
        cellElement.innerText = `${numNeighbours[i][j]}`;
        if (cell === 1) {
          cellElement.style.color = 'white';
        }
      }
      gridElement.appendChild(cellElement);
    });
  });
}

// Отрисовка сетки
function renderGrid(grid) {
  if (addNumsOnState) {
    fetch('/api/v1/neighbours')
      .then(response => response.json())
      .then(data => numNeighbours = data)
      .then(() => realRenderGrid(grid));
  } else {
    realRenderGrid(grid);
  }
}

function transformCoordsToTorus(row, col) {
  if (row == -1) {
    row = gridSizeRows - 1;
  } else if (row == gridSizeRows) {
    row = 0;
  }
  if (col == -1) {
    col = gridSizeCols - 1;
  } else if (col == gridSizeCols) {
    col = 0;
  }
  return [row, col];
}

function renderCellNumNeighbours(row, col, data) {
  let move = [-1, 0, 1];
  for (var i = 0; i < 3; i++) {
    for (var j = 0; j < 3; j++) {
      var cur_row;
      var cur_col; 
      [cur_row, cur_col] = transformCoordsToTorus(row + move[i], col + move[j]);
      var currentCell = document.getElementById(`${cur_row}_${cur_col}`); 
      currentCell.innerText = `${data[1 + move[i]][1 + move[j]]}`;
    }
  }
}

function renderCell(row, col, cell_state) {
  var currentCell = document.getElementById(`${row}_${col}`)
  if (cell_state == 1) {
    currentCell.classList.add('alive');
    currentCell.style.color = 'white';
  } else {
    currentCell.classList.remove('alive')
    currentCell.style.color = 'black';
  }  
  if (addNumsOnState) {
    fetch(`/api/v1/near_neighbours?row=${row}&col=${col}`)
    .then(response => response.json())
    .then(data => renderCellNumNeighbours(row, col, data));
  }
}

// Изменение состояния клетки
function toggleCell(row, col) {
  fetch(`/api/v1/toggle?row=${row}&col=${col}`)
    .then(response => response.json())
    .then(data => renderCell(row, col, data));
}

// Переход к следующему поколению
function nextGeneration() {
  fetch('/api/v1/next')
    .then(response => response.json())
    .then(data => renderGrid(data)); 
}

const toggleAutoButton = document.querySelector('.toggle-gen-button');

function toggleAutoGeneration() {
  autoGenerating = !autoGenerating;
  if (autoGenerating) {
    toggleAutoButton.style.color = 'green';
    autoGeneration();
  } else {
    toggleAutoButton.style.color = 'red';
  }
}

function toggleNumNeighbours() {
  addNumsOnState = !addNumsOnState;
  loadGrid();  
}

async function autoGeneration() {
  if (autoGenerating) {
    for (;;) {
      nextGeneration();
      await sleep(250);
      if (!autoGenerating) {
        break
      }
    }
  }
}

// Рандомное состояние
function seedState() {
  fetch('/api/v1/seed?fill=15')
    .then(response => response.json())
    .then(data => renderGrid(data)); 
}

function clearState() {
  fetch('/api/v1/seed?fill=0')
    .then(response => response.json())
    .then(data => renderGrid(data));
}

// Первоначальная загрузка
gridSize();
loadGrid();