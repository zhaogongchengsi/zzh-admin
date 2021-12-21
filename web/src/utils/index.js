export function JudgeRequestStatus(state) {
  if (state.success === true && state.code === 200) {
    return function (ok, err) {
      ok(state);
    };
  } else {
    return function (ok, err) {
      err(state.code, state.message);
    };
  }
}

export function isThereChinese(str) {
  const reg = /[\u4E00-\u9FA5]|[\uFE30-\uFFA0]/gi;
  if (!reg.exec(str)) {
    return false;
  } else {
    return true;
  }
}

/*
  传入一个高阶函数 第一个参数 为父亲 第二个 参数为 儿子
 */
export function DadLookSon(condition, dad, son) {
  const result = []
  dad.forEach(dad => {
    let isDad = condition(dad) // 找父节点
    if (isDad === false) {
      return false
    }
    son.forEach(sonItem => {
      let isSon = isDad(sonItem) // 找子节点
      if (isSon === false) {
        return false
      }
      let _children = DadLookSon(condition, [isSon], son) // 找子节点的子节点
      if (_children.length > 0) {
        isSon.children = _children
      }
      result.push(isSon)
    })
  });
  return result
}