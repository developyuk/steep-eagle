<template lang="pug">
  .row
    .col-md-12
      h4.title Classes
    .col-md-12.card
      .card-header
        .buttons.text-right
          router-link(to="/admin/classes/create").btn.btn-primary.btn-fill
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
            el-table-column(:key="'Image'" :min-width="100" :prop="'module.image'" :label="'Image'")
              template(slot-scope='props')
                .img-container
                  img(:src='props.row.module.image' :alt='props.row.module.name')
            el-table-column(v-for="column in tableColumns" :key="column.label" :min-width="column.minWidth" :prop="column.prop" :label="column.label" :className="column.className" :labelClassName="column.labelClassName" :sortable="column.sortable")
            <!--el-table-column(min-width='120' label="Classes")-->
            <!--template(slot-scope='props')-->
            <!--.row-->
            <!--.col-sm-3-->
            <!--.img-container-->
            <!--img(:src='props.row.module.image' alt='Image')-->
            <!--.col-sm-9-->
            <!--.module {{props.row.module.name}}-->
            <!--.branch {{props.row.branch.name}}-->
            <!--.time-->
            <!--span.day {{props.row.day}}-->
            <!--span ,&nbsp;-->
            <!--span.start {{props.row.start_at}}-->
            <!--span &nbsp;-&nbsp;-->
            <!--span.finish {{props.row.finish_at}}-->
            el-table-column(:min-width="120" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/admin/classes/${props.row.id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
                  i.ti-pencil-alt
                a.btn.btn-simple.btn-xs.btn-danger.btn-icon.remove(@click="handleDelete(props.$index, props.row)")
                  i.ti-close
        .col-sm-6.pagination-info
          p.category Showing {{from + 1}} to {{to}} of {{total}} entries
        .col-sm-6
          p-pagination.pull-right(v-model="pagination.currentPage" :per-page="pagination.perPage" :total="pagination.total")
</template>
<script>
  import Vue from 'vue'
  import {Table, TableColumn, Select, Option} from 'element-ui'
  import PPagination from 'src/components/UIComponents/Pagination.vue'
  import axios from 'axios'
  import _range from 'lodash/range'

  Vue.use(Table)
  Vue.use(TableColumn)
  Vue.use(Select)
  Vue.use(Option)
  export default {
    components: {
      PPagination
    },
    computed: {
      queriedData() {
        return this.tableData
      },
      to() {
        let highBound = this.from + this.pagination.perPage
        if (this.total < highBound) {
          highBound = this.total
        }
        return highBound
      },
      from() {
        this.getData();
        return this.pagination.perPage * (this.pagination.currentPage - 1)
      },
      total() {
        return this.pagination.total
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
        searchQuery: '',
        propsToSearch: ['start_at','finish_at'],
        tableColumns: [
//          {
//            prop: 'module.image',
//            label: 'Image',
//            minWidth: 150
//          },
          {
            prop: 'module.name',
            label: 'Module',
            minWidth: 150,
            labelClassName: 'text-capitalize',
            className: 'text-uppercase',
            sortable: true
          },
          {
            prop: 'branch.name',
            label: 'Branch',
            minWidth: 150,
            className: 'text-capitalize',
            sortable: true
          },
          {
            prop: 'day',
            label: 'Day',
            minWidth: 100,
            className: 'text-capitalize',
            sortable: true
          },
          {
            prop: 'start_at',
            label: 'Start',
            minWidth: 100,
            sortable: true
          },
          {
            prop: 'finish_at',
            label: 'Finish',
            minWidth: 100,
            sortable: true
          },
        ],
        tableData: []
      }
    },
    methods: {
      handleEdit(index, row) {
        alert(`Your want to edit ${row.module.name}`)
      },
      handleDelete(index, row) {
        let indexToDelete = this.tableData.findIndex((tableRow) => tableRow.id === row.id)
        if (indexToDelete >= 0) {
          this.tableData.splice(indexToDelete, 1)
        }
      },
      getData() {
        const config = {
          params: {
            max_results: this.pagination.perPage,
            page: this.pagination.currentPage,
          }
        };
        if (!!this.searchQuery) {
          const qList = this.propsToSearch.map(v => {
            const q = {};
            q[v] = `ilike(\"%${this.searchQuery}%\")`
            return q
          });
          config.params['where'] = {or_: qList}
        }
        axios.get(`${process.env.DBAPI}/classes_ts`, config)
          .then(response => {
            this.tableData = response.data._items;
            this.pagination.total = response.data._meta.total
            this.pagination.currentPage = response.data._meta.page
          })
          .catch(error => console.log(error, error.response))
      }
    },
    mounted() {
      this.getData();
    }
  }
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

  .branch, .time {
    text-transform: capitalize;
  }

  .time {
    min-width: 16rem;
  }
</style>
