<template>
  <div class="intro-page">
    <h1>Welcome to Solo Adventure Picker</h1>
    <p>Select your region to get started:</p>

    <select v-model="selectedRegion">
      <option disabled value="">-- Select Region --</option>
      <option v-for="region in regions" :key="region.value" :value="region.value">
        {{ region.label }}
      </option>
    </select>

    <button :disabled="!selectedRegion" @click="startAdventure">
      Start Adventure
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const selectedRegion = ref('')

const regions = [
  { value: 'bay-area', label: 'Bay Area' },
  { value: 'north-bay', label: 'North Bay' },
  { value: 'south-bay', label: 'South Bay' },
  { value: 'east-bay', label: 'East Bay' },
  // expand this as needed
]

const startAdventure = () => {
  router.push({ path: '/adventure', query: { region: selectedRegion.value } })
}
</script>

<style scoped>
.intro-page {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-width: 400px;
  margin: 0 auto;
  padding: 2rem;
  text-align: center;
}

select, button {
  padding: 0.5rem;
  font-size: 1rem;
}
</style>
