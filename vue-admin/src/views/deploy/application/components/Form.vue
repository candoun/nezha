<template>
  <div class="application-container">
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
        <el-button :size="size" type="primary" v-show="!readonly" @click.native="submitForm" >保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
	export default {
		name: 'ApplicationForm',
		props: {
      dialogVisible: {
      	type: Boolean,
      	default: true
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
      },
      dataForm: {
      	type: Object,
      	default: null
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
		},
		data() {
			return {

			}
		},
		methods: {
			handleClick: function() {
				this.$emit('click', {})
			}
		},
		mounted() {

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
