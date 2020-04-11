<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :model="Filters" :size="size">
        <el-form-item>
          <el-input v-model="Filters.name" placeholder="关键字"></el-input>
        </el-form-item>
        <el-form-item>
          <nz-button icon="fa fa-search" :label="'搜索'"  type="primary" @click="getList"/>
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
          <span>{{ scope.row.project_id }}</span>
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
          <router-link :to="'/application/edit/'+scope.row.ID">
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
import { fetchList } from '@/api/application'
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
      total: 0,
      listLoading: true,
      readonly: false,
      listQuery: {
        page: 1,
        limit: 10
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList(params) {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleDelete: function(row) {
      this.$confirm('确定删除'+ row.name +'吗？', '提示', {
        type:'warning'
      }).then(() => {
        let params = []
        let idArray = (ids+'').split(',')
        for (var i=0; i<idArray.length; i++) {
          params.push({'id': idArray[i]})
        }
        this.loading = true
        let callback = (res) => {
          if (res.code == 200) {
            this.$message({message: '删除成功', type: 'success'})
            this.list(this.pageRequest)
          } else {
            this.$message({message: '操作失败, ' + res.msg, type: 'error'})
          }
          this.loading = false
        }
        this.$api.yaml.batchDelete(params)
        .then(
          callback
        )
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
