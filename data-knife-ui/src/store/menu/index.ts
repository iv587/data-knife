import {defineStore} from 'pinia'
import {MenuItem} from '@/type'

const useMenuStore = defineStore('menu', {
    state: () => {
        return {
            activeIndex: '',
            collapse: false,
            menuList: [
                {
                    index: "/connections",
                    title: "链接管理",
                    icon: 'connection-icon',
                }] as MenuItem[]
        }
    },
    actions: {
        updateMenuItem(menuItem: MenuItem, noCreate?: boolean) {
            let notFound = true
            for (let i = 0 ; i < this.menuList.length; i ++) {
                const item = this.menuList[i]
                if (item.id == menuItem.id) {
                    item.icon = menuItem.icon
                    item.title = menuItem.title
                    item.children = menuItem.children
                    notFound = false
                    break
                }
            }
            if (noCreate) {
                return
            }
            if (notFound) {
                this.menuList.push(menuItem)
            }
        },
        setActiveIndex(index: string) {
            this.activeIndex = index
        },
        toggleCollapse() {
            this.collapse = !this.collapse
        },
        setCollapse(collapse: boolean) {
            this.collapse = collapse
        }
    }
})

export default useMenuStore