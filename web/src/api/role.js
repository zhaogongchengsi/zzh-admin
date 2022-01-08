import { Get, Post } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

export function createRole(data) {
  return new Promise((resolve, reject) => {
    Post("role/create_role", data)
      .then((response) => {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (_, msg) => {
            ElMessage.error(msg);
            reject(msg);
          }
        );
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export function getRole(lading = false) {
  return new Promise((resolve, reject) => {
    Get("role/get_roles", lading)
      .then((response) => {
        JudgeRequestStatus(response.state)(
          () => {
            resolve(response.data);
          },
          (_, msg) => {
            ElMessage.error(msg);
            reject(msg);
          }
        );
      })
      .catch((err) => {
        reject(err);
      });
  });
}
