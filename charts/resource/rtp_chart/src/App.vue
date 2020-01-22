<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <!-- <hello></hello> -->
    <charts></charts>
  </div>
</template>

<script>

import Hello from './components/Hello'
import Charts from './components/Charts'

export default {
  name: 'app',
  components:{
    Hello,
    Charts
  }
}
</script>