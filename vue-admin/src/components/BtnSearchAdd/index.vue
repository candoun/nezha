<template>
	<el-button :size="size" :type="type" :icon="icon"
		:loading="loading" @click="handleClick">
		{{label}}
	</el-button>
</template>

<script>
	export default {
		name: 'BtnSearchAdd',
		props: {
			label: {
				type: String,
				default: 'Button'
			},
			icon: {
				type: String,
				default: ''
			},
			size: {
				type: String,
				default: 'mini'
			},
			type: {
				type: String,
				default: null
			},
			loading: {
				type: Boolean,
				default: false
			},
			disabled: {
				type: Boolean,
				default: false
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
</style>
