import { Get } from "@/service/index.js";
import { ElMessage } from "element-plus";
import { JudgeRequestStatus } from "@/utils";

// import { dataToTree, copyRouter, filePathCompile } from "@/utils/asyncRoute.js";
// const routerFile = import.meta.glob("../views/**/*.{vue,js,ts,jsx,tsx}");

export function GetRouter() {
  return new Promise(function (resolve, reject) {
    Get("/menu/get_menu")
      .then(function (response) {
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
      .catch(function (error) {
        reject(error);
      });
  });
}
