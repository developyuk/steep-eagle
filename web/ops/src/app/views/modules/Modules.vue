<template lang="pug">
  .row
    .col-md-12
      h4.title Modules
    .col-md-12.card
      .card-header
        .buttons.text-right
          router-link(to="/admin/modules/create").btn.btn-primary.btn-fill
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
          el-table.table-striped(:data="tableData" border="" style="width: 100%")
            el-table-column(:key="tableColumns[0].label" :min-width="tableColumns[0].minWidth" :prop="tableColumns[0].prop" :label="tableColumns[0].label" :className="tableColumns[0].className" :sortable="tableColumns[0].sortable")
              template(slot-scope='props')
                .img-container
                  img(:src='props.row.image' :alt='props.row.name')
            el-table-column(v-for="column in tableColumns.slice(1)" :key="column.label" :min-width="column.minWidth" :prop="column.prop" :label="column.label" :className="column.className" :sortable="column.sortable")

            el-table-column(:min-width="48" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/admin/modules/${props.row._id}`").btn.btn-simple.btn-xs.btn-success.btn-icon.edit
                  i.ti-eye
                router-link(:to="`/admin/modules/${props.row._id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
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
import PPagination from "src/app/components/Pagination.vue";
import axios from "axios";
import _range from "lodash/range";
import swal from "sweetalert2";

Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Select);
Vue.use(Option);
export default {
  components: {
    PPagination
  },
  computed: {
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
          prop: "image",
          label: "Image",
          minWidth: 64
        },
        {
          prop: "name",
          label: "Name",
          minWidth: 200,
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
                .delete(`${process.env.API}/modules/${row._id}`, {
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
          sort: "-_updated",
          max_results: this.pagination.perPage,
          page: this.pagination.currentPage
        },
        headers: {
          // "cache-control": "no-cache"
        }
      };
      if (!!this.searchQuery) {
        const qList = this.propsToSearch.map(v => {
          const q = {};
          q[v] = this.searchQuery;
          return q;
        });
        config.params["where"] = { $or: qList };
      }
      axios
        .get(`${process.env.API}/modules`, config)
        .then(response => {
          this.tableData = response.data._items;
          
          const paginationTotal = this.pagination.currentPage * this.pagination.perPage;
          if (this.tableData.length == this.pagination.perPage) {
            this.pagination.total = paginationTotal + 1;
          }else{ 
            this.pagination.total = paginationTotal;
          }
        })
        .catch(error => console.log(error, error.response));
    }
  },
  mounted() {
  }
};
</script>
<style scoped lang="scss">
</style>
