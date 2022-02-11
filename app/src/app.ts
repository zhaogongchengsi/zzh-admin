import { ipcRenderer } from 'electron'

export function close (type:String) {
    ipcRenderer.send('onClose', type)
}

