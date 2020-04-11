<template>
  <div class="application-container">
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
      <el-form-item label="名称" prop="name">
        <el-input v-model="dataForm.name" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="dataForm.description" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>

      <el-form-item label="Git" prop="git">
        <el-input v-model="dataForm.git" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="Jenkins" prop="jenkins">
        <el-input v-model="dataForm.jenkins" :readonly="readonly" auto-complete="off" ></el-input>
      </el-form-item>
    </el-form>
    <div v-if="this.opType==='detail'">
      <router-link :to="'/application/list'">
      <el-button :size="size">返回</el-button>
      </router-link>
    </div>
    <div v-if="this.opType==='edit'">
      <el-button :size="size" @click.native="resetForm">Reset</el-button>
      <el-button :size="size" type="primary" v-show="!readonly" @click.native="submitForm" >保存</el-button>
    </div>
    <div v-if="this.opType==='create'">
      <el-button :size="size" @click.native="resetForm">Cancel</el-button>
      <el-button :size="size" type="primary" v-show="!readonly" @click.native="submitForm" >保存</el-button>
    </div>
  </div>
</template>

<script>
import { fetchApplication } from '@/api/application'

	export default {
		name: 'ApplicationForm',
		props: {
      id: {
      	type: Number,
      	default: 0
      },
      opType: {
      	type: String,
      	default: ''
      },
			title: {
				type: String,
				default: 'Title'
			},
			size: {
				type: String,
				default: 'mini'
			},
      readonly: {
      	type: Boolean,
      	default: false
      }
		},
		data() {
			return {
        dataForm: {
          id: 0,
          name: 'k8s-yaml',
          description: '描述',
          tag: 'dev',
          content: 'apiVersion...',
        },
        dataFormRules: {
          name: [
            {
              required: true,
              message: '请输入yaml名称',
              trigger: 'blur',
            }
          ]
        }

			}
		},
    created() {
      if (this.opType) {
        console.log("11111")
        this.getDetail(this.id)
      }
    },
		methods: {
      getDetail(params) {
        this.listLoading = true
        fetchApplication(params).then(response => {
          this.dataForm = response.data
          this.listLoading = false
        })
      },
      resetForm: function() {
        this.dialogVisible = false
        this.readonly = false
        this.$refs['dataForm'].resetFields()
      }
		}
	}
</script>

<style>
  .application-container {
    position: relative;

    .createPost-main-container {
      padding: 40px 45px 20px 50px;

      .postInfo-container {
        position: relative;
        @include clearfix;
        margin-bottom: 10px;

        .postInfo-container-item {
          float: left;
        }
      }
    }

    .word-counter {
      width: 40px;
      position: absolute;
      right: 10px;
      top: 0px;
    }
  }

  .article-textarea /deep/ {
    textarea {
      padding-right: 40px;
      resize: none;
      border: none;
      border-radius: 0px;
      border-bottom: 1px solid #bfcbd9;
    }
  }
</style>
