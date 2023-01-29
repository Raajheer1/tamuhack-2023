<template>
  <div class="flex flex-col mt-10">
    <h2 class="text-xl">Your recent searches</h2>
    <div class="flex flex-col mt-5">
      <div class="flex justify-between">
        <div class="flex">
          <i class="fa-solid fa-magnifying-glass mt-1"></i>
          <p class="pl-2">London</p>
        </div>
        <div class="bg-zinc-700 text-white text-xs rounded-full px-2 py-1">
          2 days ago
        </div>
      </div>
      <div class="flex justify-between mt-5">
        <div class="flex">
          <i class="fa-solid fa-magnifying-glass mt-1"></i>
          <p class="pl-2">Paris</p>
        </div>
        <div class="bg-zinc-700 text-white text-xs rounded-full px-2 py-1">
          3 days ago
        </div>
      </div>
      <div class="flex justify-between mt-5">
        <div class="flex">
          <i class="fa-solid fa-magnifying-glass mt-1"></i>
          <p class="pl-2">New York</p>
        </div>
        <div class="bg-zinc-700 text-white text-xs rounded-full px-2 py-1">
          4 days ago
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "RecentSearches"
};
</script>

<style scoped>

</style>