<template>
  <div class="search-bar">
    <input type="text" v-model="search" placeholder="Search items..."/>
  </div>
  <div class="list-view-container">
    <h1>Featured Items</h1>
    <ul class="item-list">
      <li v-for="(item, index) in filteredItems" :key="index" class="item">
        <h2>{{ item.title }}</h2>
        <p>{{ item.description }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'ListPage',
  data() {
    return {
      search: '',
      items: Array.from({length: 100}, (_, i) => ({
        title: `Item ${i + 1}`,
        description: `This is a description of item ${i + 1}.`
      })),
    };
  },
  computed: {
    filteredItems() {
      return this.items.filter(item => {
        return item.title.toLowerCase().includes(this.search.toLowerCase());
      });
    }
  }
};
</script>

<style scoped>
.list-view-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
}

.item-list {
  list-style: none;
  padding: 0;
}

.item {
  background-color: #cee5ac;
  border-radius: 10px;
  padding: 20px;
  margin-bottom: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.item h2 {
  margin-top: 0;
}

.item p {
  color: #666;
}
</style>