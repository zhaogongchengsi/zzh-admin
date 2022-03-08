import { confirmStatus, Post, Get } from "@/utils/service";
import { state, file } from "@/types/response";

export async function uploadFile(option: any): Promise<any> {
  const { onProgress, onError, onSuccess, fileItem, name } = option;
  const fd = new FormData();
  fd.append("file", fileItem.file);
  return new Promise<any>((resolve, reject) => {
    Post("/upload_download/upload", fd)
      .then((response) => {
        const state = confirmStatus(<state>response.data.state);
        if (state) {
          resolve(response.data.data);
          onSuccess(response);
        } else {
          onError(response);
        }
      })
      .catch((error) => {
        onError(error);
        reject(false);
      });
  });
}

export async function getfiles(type: string, limit?: number, offset?: number) :Promise<any> {
  return new Promise<any>((resolve, reject) => {
    Get(`/upload_download/get_files?type=${type}&limit=${limit}&offset=${offset}`)
      .then((response) => {
        const state = confirmStatus(<state>response.data.state);
        if (state) {
          resolve(response.data.data);
        } else {
        }
      })
      .catch(reject);
  });
}
