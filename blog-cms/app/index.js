const { app, BrowserWindow, ipcMain} = require("electron");
const path = require("path");
let win = null;
function createWindow() {
  const webPreferences = {
    nodeIntegration: true, 
    nodeIntegrationInSubFrames: true,
    nodeIntegrationInWorker: true, // 启用node
    preload: path.join(__dirname, 'preload.js')
  };
  const options = {
    width: 800,
    height: 600,
    // transparent: true, // 窗口透明
    icon: path.join(__dirname, "../public/favicon.ico"),
    // frame: false, // 创建无边框窗口
    webPreferences
  }
  const win = new BrowserWindow(options);
  // win.loadFile("../index.html");
  win.loadURL("http://localhost:3000/");
}

app.on("window-all-closed", function () {
  if (process.platform !== "darwin") app.quit();
});

app.whenReady().then(() => {
  createWindow();
  app.on("activate", function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});


// ipcMain.addListener("onClose", function (e,d) {
//   console.log("onClose",d)
// })