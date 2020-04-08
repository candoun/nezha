<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :model="Filters" :size="size">
        <el-form-item>
          <el-input v-model="Filters.name" placeholder="关键字"></el-input>
        </el-form-item>
        <el-form-item>
          <btn-search-add icon="fa fa-search" :label="'搜索'"  type="primary" @click="getList"/>
        </el-form-item>
        <el-form-item>
          <btn-search-add icon="fa fa-plus" :label="'新增'"  type="primary" @click="handleAdd"/>
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


      <el-table-column width="360px" align="center" >
        <template slot-scope="scope">
          <router-link :to="'/form/edit/'+scope.row.ID">
            <el-button plain type="success" :size="size" icon="el-icon-info">
              详情
            </el-button>
            <el-button plain type="danger" :size="size" icon="el-icon-delete">
              删除
            </el-button>
            <el-button plain type="warning" :size="size" icon="el-icon-edit">
              编辑
            </el-button>
          </router-link>

        </template>
        <template slot-scope="scope">
          <btn-search-add
            icon="fa fa-trash"
            :label="'删除'"
            :size="size"
            type="info"
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

    <el-dialog
      :title="title"
      width="80%"
      :visible="dialogVisible"
      :close-on-click-modal="false">
      <el-form
        :model="dataForm"
        label-width="100px"
        :rules="dataFormRules"
        ref="dataForm"
        :size="size"
        label-position="right">
        <el-form-item label="ID" prop="id" v-if="false">
          <el-input v-model="dataForm.id" :disabled="true" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="yaml文件名" prop="name">
          <el-input v-model="dataForm.name" :readonly="readonly" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="dataForm.description" :readonly="readonly" auto-complete="off"></el-input>
        </el-form-item>

        <el-form-item label="Tag" prop="tag">
          <el-input v-model="dataForm.tag" :readonly="readonly" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="文件内容" prop="content">
          <el-input v-model="dataForm.content" :readonly="readonly" auto-complete="off" type="textarea" :rows="20"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button :size="size" @click.native="resetForm">Cancel</el-button>
        <el-button :size="size" type="primary" v-show="!readonly" @click.native="submitForm" :loading="editLoading">保存</el-button>
      </div>
    </el-dialog>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList } from '@/api/application'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import BtnSearchAdd from '@/components/BtnSearchAdd'

export default {
  name: 'ApplicationList',
  components: {
    Pagination,
    BtnSearchAdd
  },
  filters: {
    statusFilter(state) {
      const statusMap = {
        1: 'success',
        0: 'info',
        2: 'danger'
      }
      return statusMap[state]
    },
    parseTime: parseTime
  },
  data() {
    return {
      size: 'mini',
      Filters: {
        label: ''
      },
      title: '',
      readonly: false,
      dialogVisible: false,
      list: null,
      total: 0,
      listLoading: true,
      editLoading: false,
      listQuery: {
        page: 1,
        limit: 10
      },
      dataFormRules: {
        name: [
          {
            required: true,
            message: '请输入yaml名称',
            trigger: 'blur',
          }
        ]
      },
      dataForm: {
        id: 0,
        name: 'k8s-yaml',
        description: '描述',
        tag: 'dev',
        content: 'apiVersion...',
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    // 修改 table cell边框的背景色
    tableCellStyle () {
      return 'border-color: #868686;'
    },
    // 修改 table header cell的背景色
    tableHeaderCellStyle () {
      return 'border-color: #868686; color: #606266;'
    },
    getList(params) {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    // 显示新增界面
    handleAdd: function () {
      this.dialogVisible = true
      this.title = '新建Application'
      this.readonly = false
    },
    handleEdit: function(params) {
      console.log("edit")
    },
    handleView: function(params) {
      console.log("view")
    },
    resetForm: function() {
      this.dialogVisible = false
      this.readonly = false
      this.$refs['dataForm'].resetFields()
    },
    handleDelete: function(row) {
      this.delete(row.id)
    },
    delete: function(ids) {
      this.$confirm('确定删除选中的记录吗？', '提示', {
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
