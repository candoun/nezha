<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :size="size">
        <el-form-item>
          <el-input
            v-model="listQuery.name"
            style="width: 300px;"
            placeholder="关键字"
            @keyup.enter.native="handleFilter" />
        </el-form-item>
        <el-form-item>
          <nz-button icon="el-icon-search" :label="'搜索'"  type="primary" @click="handleFilter"/>
        </el-form-item>
        <el-form-item>
          <router-link :to="'/project/create/'">
            <nz-button icon="fa fa-plus" :label="'新增'"  type="primary"/>
          </router-link>
        </el-form-item>
      </el-form>
    </div>

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row
      style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.ID }}</span>
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="应用名称">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" label="中文名">
        <template slot-scope="scope">
          <span>{{ scope.row.cn_name }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" label="简短描述">
        <template slot-scope="scope">
          <span>{{ scope.row.description }}</span>
        </template>
      </el-table-column>
      <el-table-column  width="300px" align="center" >
        <template slot-scope="scope">
          <router-link :to="'/project/update/'+scope.row.ID">
            <nz-button
              icon="el-icon-edit"
              :plain="true"
              :circle="true"
              :isLabel="false"
              :size="size"
              type="info" />
          </router-link>
          <nz-button
            icon="el-icon-delete"
            :plain="true"
            :circle="true"
            :isLabel="false"
            :size="size"
            type="danger"
            @click="handleDelete(scope.row)" />
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList, deleteProject } from '@/api/project'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import NzButton from '@/components/NzButton'
import ProjectForm from './components/form'

export default {
  name: 'ProjectList',
  components: {
    Pagination,
    NzButton,
    ProjectForm
  },

  data() {
    return {
      size: 'mini',
      Filters: {
        label: ''
      },
      title: '',
      dialogVisible: false,
      list: null,
      total: 0,
      listLoading: true,
      readonly: false,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    handleDelete: function(row) {
      this.$confirm('确定删除'+ row.name +'吗？', '提示', {
        type:'warning'
      }).then(() => {
        deleteProject(row.ID).then(response => {
          this.getList()
          if (response.msg === 'fail') {
            this.$notify({
              title: 'Fail',
              message: response.detail,
              type: 'error',
              duration: 2000
            })
          } else {
            this.$notify({
              title: 'Success',
              message: '删除用户成功!',
              type: 'success',
              duration: 2000
            })
          }
        })
      }).catch(() => {})
    },
  }
}
</script>

<style scoped>
.menu-header {
  text-align: left;
  font-size: 16px;
  color: rgb(20, 89, 121);
}
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
