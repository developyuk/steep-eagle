<template lang="pug">
  .row
    .col-md-12
      h4.title Branches
    .col-md-12.card
      .card-header
        .buttons.text-right
          router-link(to="/admin/branches/create").btn.btn-primary.btn-fill
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
            el-table-column(:min-width="32" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/admin/branches/${props.row._id}`").btn.btn-simple.btn-xs.btn-success.btn-icon.edit
                  i.ti-eye
                router-link(:to="`/admin/branches/${props.row._id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
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
import swal from "sweetalert2";

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
          // className: "text-capitalize",
          sortable: true
        }
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
        swal({
          // title: "Input something",
          title: "Are you sure?",
          html: `You won't be able to revert this!.
          Submit <strong>${row.name}</strong> below`,
          type: "warning",
          input: "text",
          showCancelButton: true,
          showLoaderOnConfirm: true,
          confirmButtonClass: "btn btn-success btn-fill",
          cancelButtonClass: "btn btn-danger btn-fill",
          buttonsStyling: false,
          preConfirm: text => {
            if (text === row.name) {
              return axios
                .delete(`${process.env.API}/branches/${row._id}`, {
                  headers: { "if-match": row._etag }
                })
                .then(response => {
                  // this.tableData.splice(indexToDelete, 1);
                  this.getData();
                  return {
                    title: "Deleted!",
                    type: "success",
                    text: `<strong>${text}</strong> has been deleted.`
                  };
                })
                .catch(error => {
                  swal.showValidationError(`Request failed: ${error}`);
                });
            } else {
              return new Promise((resolve, reject) => {
                resolve({
                  title: "Error",
                  type: "error",
                  text: `Please type <strong>${row.name}</strong>`
                });
              });
            }
          },
          allowOutsideClick: () => !swal.isLoading()
        })
          .then(result => {
            console.log(result);
            swal({
              title: result.title,
              html: result.text,
              type: result.type,
              timer: 1000,
              confirmButtonClass: "btn btn-success btn-fill",
              buttonsStyling: false
            });
          })
          .catch(swal.noop);
      }
    },
    getData() {
      const config = {
        params: {
          max_results: this.pagination.perPage,
          page: this.pagination.currentPage,
          sort: "-_updated"
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
        .get(`${process.env.API}/branches`, config)
        .then(response => {
          this.tableData = response.data._items;
          this.pagination.currentPage = response.data._meta.page;
          if (this.tableData.length == this.pagination.perPage) {
            this.pagination.total =
              this.pagination.currentPage * this.pagination.perPage + 1;
          }
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
