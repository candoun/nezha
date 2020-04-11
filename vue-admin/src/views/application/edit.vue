<template>
  <application-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'edit'" />
</template>

<script>
import ApplicationForm from './components/form'

export default {
  name: 'EditApplication',
  components: { ApplicationForm }
}
</script>
