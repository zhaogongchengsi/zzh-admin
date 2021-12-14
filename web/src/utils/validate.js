import { isThereChinese } from "./index";

/*       
          parent_id: "",
          component: "",
          menu_path: "",
          menu_name: "",
          menu_icon: "",
          menu_disable: false,
          Label: "",
          sort: "",
          remarks: "" 
        */

function valiStr(rule, value, callback) {
  if (typeof value !== "string") {
    callback(new Error("请输入字符串"));
  } else if (isThereChinese(value)) {
    callback(new Error("输入的值内包含中文字符,请输入英文字符"));
  }
}

function valiNum(rule, value, callback) {
  let _va = parseInt(value);
  if (Number.isNaN(_va)) {
    callback(new Error("请输入数字"));
  }
}

export const menuRules = {
  parent_id: [
    {
      required: true,
      message: "请输入父节点",
      trigger: "change",
    },
    {
      validator: valiNum,
      trigger: "blur",
    },
  ],
  component: [
    {
      required: true,
      message: "请输入组件的路径",
      trigger: "change",
    },
    {
      validator: valiStr,
      trigger: "change",
    },
  ],
  menu_path: [
    {
      required: true,
      message: "请输入路由的路径",
      trigger: "change",
    },
    {
      validator: valiStr,
      trigger: "change",
    },
  ],
  menu_disable: [
    {
      required: true,
      message: "请选择是否禁用",
      trigger: "blur",
    },
  ],
  Label: [
    {
      required: true,
      message: "请输入展示的名称",
      trigger: "blur",
    },
  ],
  sort: [
    {
      required: true,
      message: "请输入排序标记",
      trigger: "blur",
    },
    {
      validator: valiNum,
      trigger: "blur",
    },
  ],
};
