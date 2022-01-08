import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

export function createMenu(data) {
  return new Promise(function (resolve, reject) {
    Post("/menu/create_menu", data)
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (code, msg) => {
            ElMessage.error(msg);
            reject({});
          }
        );
      })
      .catch(function (err) {
        reject({});
      });
  });
}

export function findMenuId(id, lading = false) {
  return new Promise(function (resolve, reject) {
    Post(
      "/menu/find_menu",
      {
        id,
      },
      lading
    )
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (code, msg) => {
            ElMessage.error(msg);
            reject({});
          }
        );
      })
      .catch(function (err) {
        reject({});
      });
  });
}

export function upMenu(menu) {
  return new Promise(function (resolve, reject) {
    Post("/menu/up_menu", menu, true)
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(true);
          },
          (code, message) => {
            ElMessage.error(message);
            reject({});
          }
        );
      })
      .catch(function () {
        reject({});
      });
  });
}

export function getSubMenus(id) {
  return new Promise((resolve, reject) => {
    Post("/menu/get_submenu", { id }, false)
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (code, msg) => {
            ElMessage.error(msg);
            reject({});
          }
        );
      })
      .catch(function (err) {
        reject({ err });
      });
  });
}

export function deleteMenu(id) {
  return new Promise((resolve, reject) => {
    Post("/menu/delete_menu", { id }, false)
      .then(function (response) {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(true);
          },
          () => {
            reject({});
          }
        );
      })
      .catch(function (err) {
        reject({ err });
      });
  });
}
