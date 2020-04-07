<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
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


      <el-table-column align="center" label="--" >
        <template slot-scope="scope">
          <router-link :to="'/form/edit/'+scope.row.ID">
            <el-button type="success" size="small" icon="el-icon-info">
              详情
            </el-button>
            <el-button type="danger" size="small" icon="el-icon-delete">
              删除
            </el-button>
            <el-button type="warning" size="small" icon="el-icon-edit">
              编辑
            </el-button>
          </router-link>
        </template>

      </el-table-column>
      <el-table-column align="center" label="--" >
        <template slot-scope="scope">
          <router-link :to="'/form/edit/'+scope.row.ID">
            <el-button type="primary" size="small" icon="el-icon-caret-right">
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
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ApplicationList',
  components: { Pagination },
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
      list: null,
      total: 0,
      listLoading: true,
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
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
