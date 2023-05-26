<template>
    <div>This is the Users View</div>
    <button @click="getUsers">Load Users</button>
    <li v-for="user in users">
        {{ user.name }}
    </li>
</template>
<script>
export default {
    data() {
        return {
            items: [{ message: 'Foo' }, { message: 'Bar' }],
            users: [],
        }
    },
    methods: {
        async getUsers() {
            const response = await fetch("http://localhost:5050/shepherd/api/v1/users");
            this.users = await response.json();
            console.log(this.users);
        }
    }
}
</script>
