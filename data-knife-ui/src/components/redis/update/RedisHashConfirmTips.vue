<template>
  <el-form label-width="auto" label-position="right">
    <el-form-item>
      <template #label>
        <span style="color: var(--el-color-info)">field:</span>
      </template>
      <p class="show-value" style="color: var(--el-color-primary)">{{field}}</p>
    </el-form-item>
    <el-form-item>
      <template #label>
        <span style="color: var(--el-color-info)">new value:</span>
      </template>
      <p class="show-value">{{val}}</p>
    </el-form-item>
    <el-form-item v-if="oVal != null && oVal != ''" label="old value:">
      <template #label>
        <span style="color: var(--el-color-info)">old value:</span>
      </template>
      <p style="color: var(--el-color-danger)" class="show-value">{{oVal}}</p>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
defineProps<{
  field: string,
  val: string,
  oVal?: string
}>()
</script>

<style lang="scss" scoped>
.show-value {
  width: 100%;
  color: var(--el-color-success);
  word-break: break-all;
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-base);
  padding: 5px;
}
</style>