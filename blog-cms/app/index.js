const { app, BrowserWindow, ipcMain } = require("electron");
const path = require("path");
let win = null;
function createWindow() {
  const webPreferences = {
    nodeIntegration: true,
    nodeIntegrationInSubFrames: true,
    nodeIntegrationInWorker: true, // 启用node
    preload: path.join(__dirname, "preload.js"),
  };
  const options = {
    width: 1500,
    height: 800,
    // transparent: true, // 窗口透明
    icon: path.join(__dirname, "../public/favicon.ico"),
    // frame: false, // 创建无边框窗口
    webPreferences,
  };
  win = new BrowserWindow(options);
  // win.loadFile("../index.html");
  win.loadURL("http://localhost:3000/");
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
  win.setFullScreen(true);
});

ipcMain.addListener("onFullScreenExit", function () {
  win.setFullScreen(false);
});
