const { contextBridge } = require("electron");
const { ipcRenderer } = require("electron");

function sendMessage(message, data) {
  ipcRenderer.send(message, data);
}

contextBridge.exposeInMainWorld("blogApp", {
  sendMessage,
});
