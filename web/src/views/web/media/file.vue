<script setup>
import { ref } from 'vue'

const handleChange = (file, list) => {
  console.log(file)
}
</script>


<template>
    <div class="file-container">
        <el-upload
                class="upload-demo"
                action="http://127.0.0.1:8000/api/v1/upload_download/upload"
                :on-change="handleChange"
            >
                <el-button type="primary">选择文件上传</el-button>
                <template #tip>
                    <div class="el-upload__tip">
                        选择文件上传
                    </div>
                </template>
        </el-upload>
        
    </div>
</template>


<style lang="scss" scoped>

.file-container {
    padding: 15px;
    box-sizing: border-box;
}

</style>