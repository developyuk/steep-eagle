<template lang="pug">
  .row
    .col-md-12
      h4.title Tutors
    .col-md-12.card
      .card-header
        .buttons.text-right
          router-link(to="/admin/tutors/create").btn.btn-primary.btn-fill
            span.btn-label: i.fa.fa-plus
            span Create
      .card-content.row
        .col-sm-6
          el-select.select-default(v-model="pagination.perPage" placeholder="Per page")
            el-option.select-default(v-for="item in pagination.perPageOptions" :key="item" :label="item" :value="item")
        .col-sm-6
          .pull-right
            label
              input.form-control.input-sm(type="search" placeholder="Search records" v-model="searchQuery" aria-controls="datatables")
        .col-sm-12
          el-table.table-striped(:data="queriedData" border="" style="width: 100%")
            el-table-column(v-for="column in tableColumns" :key="column.label" :min-width="column.minWidth" :prop="column.prop" :label="column.label" :className="column.className" :sortable="column.sortable")
            el-table-column(:min-width="120" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/admin/tutors/${props.row.id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
                  i.ti-pencil-alt
                a.btn.btn-simple.btn-xs.btn-danger.btn-icon.remove(@click="handleDelete(props.$index, props.row)")
                  i.ti-close
        .col-sm-6.pagination-info
          p.category Showing {{from + 1}} to {{to}} of {{total}} entries
        .col-sm-6
          p-pagination.pull-right(v-model="pagination.currentPage" :per-page="pagination.perPage" :total="pagination.total")
</template>
<script>
import Vue from "vue";
import { Table, TableColumn, Select, Option } from "element-ui";
import PPagination from "src/components/UIComponents/Pagination.vue";
import axios from "axios";
import _range from "lodash/range";
import mixinNotify from "src/app/mixins/notify";

Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Select);
Vue.use(Option);
export default {
  mixins: [mixinNotify],
  components: {
    PPagination
  },
  computed: {
    queriedData() {
      return this.tableData;
    },
    to() {
      let highBound = this.from + this.pagination.perPage;
      if (this.total < highBound) {
        highBound = this.total;
      }
      return highBound;
    },
    from() {
      this.getData();
      return this.pagination.perPage * (this.pagination.currentPage - 1);
    },
    total() {
      return this.pagination.total;
    }
  },
  data() {
    return {
      pagination: {
        perPage: 5,
        currentPage: 1,
        perPageOptions: [5, 10, 25, 50],
        total: 0
      },
      searchQuery: "",
      propsToSearch: ["name"],
      tableColumns: [
        {
          prop: "name",
          label: "Name",
          minWidth: 200,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "username",
          label: "Username",
          minWidth: 200,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "email",
          label: "Email",
          minWidth: 200,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "photo",
          label: "Photo",
          minWidth: 200,
          className: "text-capitalize",
          sortable: true
        },
      ],
      tableData: []
    };
  },
  methods: {
    handleEdit(index, row) {
      alert(`Your want to edit ${row.name}`);
    },
    handleDelete(index, row) {
      let indexToDelete = this.tableData.findIndex(
        tableRow => tableRow.id === row.id
      );
      if (indexToDelete >= 0) {
        this.tableData.splice(indexToDelete, 1);
        axios
          .delete(`${process.env.DBAPI}/users/${row.id}`, {
            headers: { "if-match": row._etag }
          })
          .then(response => {
            this.notifyVue({
              component: {
                template: `<span>Success deleted</span>`
              },
              type: "success"
            });
            this.getData();
          })
          .catch(error => {
            this.notifyVue({
              component: {
                template: `<span>Fail deleted</span>`
              },
              type: "danger"
            });
          });
      }
    },
    getData() {
      const config = {
        params: {
          max_results: this.pagination.perPage,
          page: this.pagination.currentPage,
          sort: '-_updated',
          where: {role:'tutor'}
        },
        headers: {
          "cache-control": "no-cache"
        }
      };
      if (!!this.searchQuery) {
        const qList = this.propsToSearch.map(v => {
          const q = {};
          q[v] = `ilike(\"%${this.searchQuery}%\")`;
          return q;
        });
        config.params["where"] = { or_: qList };
      }
      axios
        .get(`${process.env.DBAPI}/users`, config)
        .then(response => {
          this.tableData = response.data._items;
          this.pagination.total = response.data._meta.total;
          this.pagination.currentPage = response.data._meta.page;
        })
        .catch(error => console.log(error, error.response));
    }
  },
  mounted() {
    this.getData();
  }
};
</script>
<style>
</style>
