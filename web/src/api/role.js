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

export function getRole(lading = false, toTree = true) {
  return new Promise((resolve, reject) => {
    Get("role/get_roles", lading)
      .then((response) => {
        JudgeRequestStatus(response.state)(
          () => {
            if (toTree === true) {
              resolve(datahandler(response.data));
            } else {
              resolve(response.data);
            }
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

  function datahandler(data) {
    let _data = JSON.parse(JSON.stringify(data));

    let parArr = _data.filter((pItem) => {
      return pItem.ParentID === 0;
    });
    let sonArr = _data.filter((pItem) => {
      return pItem.ParentID !== 0;
    });

    getChildren(parArr, sonArr);

    function getChildren(data, sonArr) {
      data.forEach(function (pItem, index) {
        sonArr.forEach(function (sItem) {
          if (pItem.AuthRoleID === sItem.ParentID) {
            getChildren([sItem], sonArr);
            if (pItem.children && pItem.children.length > 0) {
              pItem.children.push(sItem);
            } else {
              pItem.children = [sItem];
            }
          }
        });
      });
    }

    return parArr;
  }
}
