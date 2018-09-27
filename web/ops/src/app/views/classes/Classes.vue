<template lang="pug">
  .row
    .col-md-12
      h4.title Classes
    .col-md-12.card
      .card-header
        .buttons
          .row
            .col-sm-6
              .btn-group
                router-link(to="/schedules").btn.btn-primary.btn-icon.btn-fill
                  i.fa.fa-list-ul
                router-link(to="/schedules/calendar").btn.btn-primary.btn-icon
                  i.fa.fa-th-large
            .col-sm-6.text-right
              router-link(to="/schedules/create").btn.btn-primary.btn-fill
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
            //- el-table-column(:key="'Id'" :min-width="64" :prop="'id'" :label="'#'")
            el-table-column(:key="'Image'" :min-width="100" :prop="'module.image'" :label="'Image'")
              template(slot-scope='props')
                .img-container
                  img(:src='props.row.module.image' :alt='props.row.module.name')
            el-table-column(v-for="column in tableColumns.slice(1)" :key="column.label" :min-width="column.minWidth" :prop="column.prop" :label="column.label" :className="column.className" :labelClassName="column.labelClassName" :sortable="column.sortable")
            el-table-column(:min-width="72" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/schedules/${props.row._id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
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
      propsToSearch: ["q"],
      tableColumns: [
        {
          prop: "id",
          label: "Id",
          minWidth: 50
        },
        {
          prop: "module.name",
          label: "Module",
          minWidth: 150,
          labelClassName: "text-capitalize",
          className: "text-uppercase",
          sortable: true
        },
        {
          prop: "branch.name",
          label: "Branch",
          minWidth: 150,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "tutor.name",
          label: "Tutor",
          minWidth: 150,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "day",
          label: "Day",
          minWidth: 100,
          className: "text-capitalize",
          sortable: true
        },
        {
          prop: "startAt",
          label: "Start",
          minWidth: 100,
          sortable: true
        },
        {
          prop: "finishAt",
          label: "Finish",
          minWidth: 100,
          sortable: true
        },
        {
          prop: "studentsTotal",
          label: "Total Students",
          minWidth: 128 + 16,
          sortable: true
        }
      ],
      tableData: []
    };
  },
  methods: {
    handleEdit(index, row) {
      alert(`Your want to edit ${row.module.name}`);
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
          Submit <strong>${row.module.name}</strong> below`,
          type: "warning",
          input: "text",
          showCancelButton: true,
          showLoaderOnConfirm: true,
          confirmButtonClass: "btn btn-success btn-fill",
          cancelButtonClass: "btn btn-danger btn-fill",
          buttonsStyling: false,
          preConfirm: text => {
            if (text === row.module.name) {
              return axios
                .delete(`${process.env.API}/classes/${row._id}`, {
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
                  text: `Please type <strong>${row.module.name}</strong>`
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
          page: this.pagination.currentPage,
          embedded: {
            branch: true,
            module: true,
            tutor: true
          }
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
        .get(`${process.env.API}/classes`, config)
        .then(response => {
          this.tableData = response.data._items;
          this.tableData = this.tableData.map(v => {
            v.students_sum = v.students.length;
            return v;
          });
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
<style scoped lang="scss">
.img-container {
  $size: 5rem;
  width: $size;
  height: $size;
  img {
    border-radius: 1rem;
  }
}

.module {
  font-size: 2rem;
  text-transform: uppercase;
}

.branch,
.time {
  text-transform: capitalize;
}

.time {
  min-width: 16rem;
}
</style>
