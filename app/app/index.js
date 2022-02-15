const { app, BrowserWindow, ipcMain } = require("electron");
const path = require("path");

let win = null;
function env() {
  if (process.env.NODE_ENV === "production" /* 生产环境 */) {
    return false;
  } else if (process.env.NODE_ENV === "development" /* 开发环境 */) {
    return true;
  }
  return false;
}

function createWindow() {
  const webPreferences = {
    nodeIntegration: true,
    nodeIntegrationInSubFrames: true,
    nodeIntegrationInWorker: true, // 启用node
    preload: path.join(__dirname, "preload.js"),
    devTools: env(),
  };
  const options = {
    width: 1500,
    height: 800,
    minWidth: 1200,
    minHeight: 800,
    transparent: !env(), // 窗口透明
    icon: path.join(__dirname, "../public/favicon1.ico"),
    frame: !env(), // 创建无边框窗口
    webPreferences,
  };
  win = new BrowserWindow(options);
  if (!env()) {
    win.loadURL(path.join(__dirname, "dist/index.html")); // 构件时 使用成路径
  } else {
    win.loadURL("http://localhost:3000/"); // 开发环境用这个
    win.webContents.openDevTools();
  }

  return win;
}

app.on("window-all-closed", function () {
  if (process.platform !== "darwin") app.quit();
  win = null;
});

app.whenReady().then(() => {
  createWindow();
  app.on("activate", function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

ipcMain.addListener("onClose", function (e, d) {
  win.close();
});

ipcMain.addListener("onMinimize", function (e, d) {
  win.minimize();
});

ipcMain.addListener("onFullScreen", function () {
  win.maximize();
});

ipcMain.addListener("onFullScreenExit", function () {
  win.unmaximize();
});
