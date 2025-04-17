<template>
  <div id="app">
    <h1>Solo Adventure Picker</h1>
    <button @click="getRandomAdventure">Pick an Adventure</button>
    <div v-if="adventure" class="card">
      <h2>{{ adventure }}</h2>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const adventure = ref('')

const getRandomAdventure = async () => {
  const res = await fetch('http://localhost:8080/random')
  const data = await res.json()
  adventure.value = data.Name
}
</script>
