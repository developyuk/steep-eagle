<template lang="pug">
  .row
    .col-md-12
      h4.title Tutors Activities
    .col-md-12.card
      .card-header
        .buttons
          .row
            .col-sm-6
              .btn-group
                router-link(to="/activities/tutors").btn.btn-primary.btn-icon.btn-fill
                  i.fa.fa-list-ul
                router-link(to="/activities/tutors/timeline").btn.btn-primary.btn-icon
                  i.fa.fa-th-large
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
            el-table-column(:key="'Id'" :min-width="64" :prop="'id'" :label="'#'")
            el-table-column(:key="'Image'" :min-width="100" :prop="'module.image'" :label="'Image'")
              template(slot-scope='props')
                .img-container
                  img(:src='props.row.attendance.module.image' :alt='props.row.attendance.module.name')
            el-table-column(v-for="column in tableColumns" :key="column.label" :min-width="column.minWidth" :prop="column.prop" :label="column.label" :className="column.className" :labelClassName="column.labelClassName" :sortable="column.sortable")
            el-table-column(:min-width="72" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/activities/tutors/${props.row.id}`").btn.btn-simple.btn-xs.btn-success.btn-icon.edit
                  i.ti-eye
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
      propsToSearch: [
        "student.name",
        "attendance.module.name",
        "attendance.class.branch.name",
        "tutor.name"
      ],
      tableColumns: [
        {
          prop: "tutor.name",
          label: "Tutor",
          minWidth: 128,
          // className: "text-capitalize",
          sortable: true
        },
        {
          prop: "attendance.module.name",
          label: "Module",
          minWidth: 192,
          // labelClassName: "text-capitalize",
          // className: "text-uppercase",
          sortable: true
        },
        {
          prop: "attendance.class.branch.name",
          label: "Branch",
          minWidth: 192,
          // className: "text-capitalize",
          sortable: true
        },
        {
          prop: "_created",
          label: "Date",
          minWidth: 192,
          sortable: true
        }
      ],
      tableData: []
    };
  },
  methods: {
    // handleEdit(index, row) {
    //   alert(`Your want to edit ${row.module.name}`);
    // },
    // handleDelete(index, row) {
    //   let indexToDelete = this.tableData.findIndex(
    //     tableRow => tableRow.id === row.id
    //   );
    //   if (indexToDelete >= 0) {
    //     swal({
    //       // title: "Input something",
    //       title: "Are you sure?",
    //       html: `You won't be able to revert this!.
    //       Submit <strong>${row.module.name}</strong> below`,
    //       type: "warning",
    //       input: "text",
    //       showCancelButton: true,
    //       showLoaderOnConfirm: true,
    //       confirmButtonClass: "btn btn-success btn-fill",
    //       cancelButtonClass: "btn btn-danger btn-fill",
    //       buttonsStyling: false,
    //       preConfirm: text => {
    //         if (text === row.module.name) {
    //           return axios
    //             .delete(`${process.env.API}/classes/${row.id}`, {
    //               headers: { "if-match": row._etag }
    //             })
    //             .then(response => {
    //               // this.tableData.splice(indexToDelete, 1);
    //               this.getData();
    //               return {
    //                 title: "Deleted!",
    //                 type: "success",
    //                 text: `<strong>${text}</strong> has been deleted.`
    //               };
    //             })
    //             .catch(error => {
    //               swal.showValidationError(`Request failed: ${error}`);
    //             });
    //         } else {
    //           return new Promise((resolve, reject) => {
    //             resolve({
    //               title: "Error",
    //               type: "error",
    //               text: `Please type <strong>${row.module.name}</strong>`
    //             });
    //           });
    //         }
    //       },
    //       allowOutsideClick: () => !swal.isLoading()
    //     })
    //       .then(result => {
    //         console.log(result);
    //         swal({
    //           title: result.title,
    //           html: result.text,
    //           type: result.type,
    //           timer: 1000,
    //           confirmButtonClass: "btn btn-success btn-fill",
    //           buttonsStyling: false
    //         });
    //       })
    //       .catch(swal.noop);
    //   }
    // },
    getData() {
      const config = {
        params: {
          sort: "-_created",
          max_results: this.pagination.perPage,
          page: this.pagination.currentPage,
          // embedded: {
          //   attendance: true,
          //   "attendance.class": true,
          //   "attendance.class.branch": true,
          //   "attendance.module": true,
          //   tutor: true
          // }
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
        .get(`${process.env.API}/attendances_tutors`, config)
        .then(response => {
          this.tableData = response.data._items;
          this.tableData.forEach(v => {
            
            v.attendance = {
              _id:v.attendance,
              module:{
                _id:0,
                name:'',
                image:'',
              },
              class:{
                _id:0,
                name:'',
                branch:{
                  _id:0,
                  name:''
                }
              },
            }
            axios
              .get(`${process.env.API}/attendances/${v.attendance._id}`)
              .then(response => {
                v.attendance = response.data;

                v.attendance.module = {
                  _id:v.attendance.module,
                  name:'',
                  image:'',
                };
                v.attendance.class = {
                  _id:v.attendance.class,
                  name:'',
                  branch:{
                    _id:0,
                    name:''
                  }
                };
                axios
                  .get(`${process.env.API}/modules/${v.attendance.module._id}`)
                  .then(response => v.attendance.module = response.data)

                axios
                  .get(`${process.env.API}/classes/${v.attendance.class._id}`)
                  .then(response => {
                    v.attendance.class = response.data;
                    v.attendance.class.branch = {
                      _id:v.attendance.class.branch,
                      name:'',
                    };

                    axios
                      .get(`${process.env.API}/branches/${v.attendance.class.branch._id}`)
                      .then(response => v.attendance.class.branch = response.data)
                  })
              })
              .catch(error => console.log(error, error.response));

            axios
              .get(`${process.env.API}/users/${v.tutor}`)
              .then(response => v.tutor = response.data)
              .catch(error => v.tutor = error.response.data);
          });

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
