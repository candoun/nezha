<template>
  <application-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'create'" />
</template>

<script>
import ApplicationForm from './components/form'

export default {
  name: 'CreateApplication',
  components: { ApplicationForm }
}
</script>
