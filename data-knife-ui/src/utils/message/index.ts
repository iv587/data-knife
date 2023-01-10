import {ElMessage} from 'element-plus'
import {ElMessageBox} from 'element-plus'
import {ElMessageBoxOptions} from "element-plus/es/components/message-box/src/message-box.type";

class MsgUtils {
    success(message?: string) {
        return new Promise<void>(resolve => {
            ElMessage.success({
                message: message,
                duration: 1000,
                center: true,
                onClose: () => {
                    resolve()
                }
            })
        })
    }

    error(message?: string) {
        return new Promise<void>((resolve, reject) => {
            ElMessage.error({
                message, duration: 2000, onClose: () => {
                    resolve()
                }
            })
        })
    }

    async confirm(title: string, message: string, options?: ElMessageBoxOptions) {
        return ElMessageBox.confirm(message, title, options)
    }
}

export default new MsgUtils()