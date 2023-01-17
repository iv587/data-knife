import {defineStore} from 'pinia'
import {MenuItem} from '@/type'
import exp from "constants";

function isMobile() {
    let flag = navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i);
    return flag
}

function addAll(list: MenuItem[], item: MenuItem) {
    list.push(item)
    if (item.children && item.children.length > 0) {
        item.children.forEach(it => {
            addAll(list, it)
        })
    }
}

const useMenuStore = defineStore('menu', {
    state: () => {
        const mobile = !!isMobile()
        return {
            activeIndex: '',
            collapse: mobile,
            menuList: [
                {
                    index: "/connections",
                    title: "链接管理",
                    icon: 'connection-icon',
                }] as MenuItem[],
            menuAllList: [{
                index: "/connections",
                title: "链接管理",
                icon: 'connection-icon',
            }] as MenuItem[]
        }
    },
    getters: {
        getActiveMenuItem(): MenuItem | undefined{
            return this.menuAllList.find(item => item.index == this.activeIndex)
        }
    },
    actions: {
        updateMenuItem(menuItem: MenuItem, noCreate?: boolean) {
            let notFound = true
            for (let i = 0; i < this.menuList.length; i++) {
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
                addAll(this.menuAllList, menuItem)
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