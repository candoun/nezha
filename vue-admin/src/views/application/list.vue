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
          <router-link :to="'/application/create/'">
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

      <el-table-column width="120px" align="center" label="所属项目">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="listProject[scope.row.project_id]['cn_name']" placement="top">
            <span>{{ listProject[scope.row.project_id]['name'] }}</span>
          </el-tooltip>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" label="Git地址">
        <template slot-scope="scope">
          <span>{{ scope.row.git }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" label="Job名称">
        <template slot-scope="scope">
          <span>{{ scope.row.jenkins }}</span>
        </template>
      </el-table-column>
      <el-table-column  width="300px" align="center" >
        <template slot-scope="scope">
          <router-link :to="'/application/detail/'+scope.row.ID">
          <nz-button
            icon="el-icon-info"
            :size="size"
            :plain="true"
            :circle="true"
            :isLabel="false"
            type="info" />
          </router-link>
          <router-link :to="'/application/update/'+scope.row.ID">
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
      <el-table-column width="120px" align="center" >
        <template slot-scope="scope">
          <router-link :to="'/form/edit/'+scope.row.ID">
            <el-button type="primary" :size="size" icon="el-icon-caret-right">
              部署
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList, deleteApplication } from '@/api/application'
import { fetchList as fetchProjectList } from '@/api/project'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import NzButton from '@/components/NzButton'
import ApplicationForm from './components/form'

export default {
  name: 'ApplicationList',
  components: {
    Pagination,
    NzButton,
    ApplicationForm
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
      listProject: null,
      total: 0,
      listLoading: true,
      readonly: false,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined
      },
      listProjectQuery: {
        page: 1,
        limit: 100,
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
      fetchProjectList(this.listProjectQuery).then(response => {
        this.listProject = response.data.list
        this.listLoading = false
      }).then(() => {
        fetchList(this.listQuery).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.listLoading = false
        })
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
        deleteApplication(row.ID).then(response => {
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
