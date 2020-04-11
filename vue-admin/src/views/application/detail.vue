<template>
  <application-form  :id=parseInt(this.$route.params.id) :readonly="true" :opType="'detail'" />
</template>

<script>
import { fetchApplication } from '@/api/application'
import ApplicationForm from './components/form'

export default {
  name: 'DetailApplication',
  components: { ApplicationForm }
}
</script>
