<script>
import { h } from 'vue'
import { ElMenu, ElSubMenu, ElMenuItem } from 'element-plus'
export default {
  name: "BaseMenus",
  props: {
    menulist: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  render () {

    function renderContextText (context) {
      return h("span", {/* props */}, context.label)
    }

    function menuItem(menu) {
      return h(ElMenuItem,{/* props */},renderContextText(menu))
    }

    const menus = this.menulist.map(function (menu, index) {
      if (menu.children && menu.children.length > 0) {
        return h(ElSubMenu,{/* props */}, menu.children.map((child, i) => {
          menuItem(child)
        }, this))
      } else {
        return menuItem(menu)
      }
    })

    return h(
      ElMenu,
      {/* props */},
      menus
    )
  }
}
</script>

