import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

export function createMenu(data) {
  return new Promise(function (resolve, reject) {
    Post("/menu/create_menu", data).then(function (response) {
      JudgeRequestStatus(response.state)(
        () => {
          resolve(response.data);
        },
        () => {
          ElMessage.error(response.message);
          reject({});
        }
      );
    });
  });
}

export function findMenuId (id) {
  return new Promise(function (resolve, reject) {
    Post("/menu/find_menu", {
      id
    }).then(function (response) {
      JudgeRequestStatus(response.state)(
        () => {
          resolve(response.data);
        },
        () => {
          ElMessage.error(response.message);
          reject({});
        }
      );
    })
  })
}
