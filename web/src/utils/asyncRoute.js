

export function dataToTree(data) {
    let parents = data.filter((p) => p.ParentId === p.ID),
      children = data.filter((c) => c.ParentId !== c.ID);
    let _children = JSON.parse(JSON.stringify(children));
    function pasParAndChildren(par, chi) {
      par.map((pitem, i) => {
        chi.map((citem, ci) => {
          if (citem.ParentId === pitem.ID) {
            _children.splice(ci, 1);
            pasParAndChildren([citem], _children);
            if (pitem.children) {
              pitem.children.push(citem);
            } else {
              pitem.children = [citem];
            }
          }
        });
      });
    }
    pasParAndChildren(parents, children);
    return parents;
  }
  
  export function copyRouter(data) {
    return data.map((item, i) => {
      if (item.children) {
        return {
          name: item.Name,
          component: item.Component,
          path: item.Path,
          label: item.Label,
          children: [...copyRouter(item.children)],
        };
      } else {
        return {
          name: item.Name,
          component: item.Component,
          path: item.Path,
          label: item.Label,
        };
      }
    });
  }



  //  替换文件路径
export function filePathCompile(files) {
    let newPaths = {};
    for (const file in files) {
      const pathreg = /\.\/([^.]+).*/i;
      const path = file.match(pathreg);
      newPaths[path[1]] = files[file];
    }
    return newPaths;
  }
  
export function filePathCompileMap(files) {
    let newPaths = new Map();
    for (const file in files) {
      const pathreg = /\.\/([^.]+).*/i;
      const path = file.match(pathreg);
      newPaths.set(path[1], files[file]);
    }
    return newPaths;
}
  
export function filePathCompileArr(files) {
    let newPaths = [];
    for (const file in files) {
      const pathreg = /\.\/([^.]+).*/i;
      const path = file.match(pathreg);
      newPaths.push({
        pathName: path[1],
        component: files[file],
      });
    }
    return newPaths;
}