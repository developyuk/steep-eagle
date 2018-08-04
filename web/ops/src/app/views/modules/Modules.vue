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
          el-table.table-striped(:data="queriedData" border="" style="width: 100%")
            el-table-column(:key="tableColumns[0].label" :min-width="tableColumns[0].minWidth" :prop="tableColumns[0].prop" :label="tableColumns[0].label" :className="tableColumns[0].className" :sortable="tableColumns[0].sortable")
              template(slot-scope='props')
                .img-container
                  img(:src='props.row.image' :alt='props.row.name')
            el-table-column(:key="tableColumns[1].label" :min-width="tableColumns[1].minWidth" :prop="tableColumns[1].prop" :label="tableColumns[1].label"  :className="tableColumns[1].className" :sortable="tableColumns[1].sortable" :labelClassName="tableColumns[1].labelClassName")

            el-table-column(:min-width="120" fixed="right" label="Actions")
              template(slot-scope="props")
                router-link(:to="`/admin/modules/${props.row.id}/edit`").btn.btn-simple.btn-xs.btn-warning.btn-icon.edit
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

Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Select);
Vue.use(Option);
export default {
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
          prop: "image",
          label: "Image",
          minWidth: 60
        },
        {
          prop: "name",
          label: "Name",
          minWidth: 200,
          className: "text-uppercase",
          labelClassName: "text-capitalize",
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
        this.tableData.splice(indexToDelete, 1);

        axios
          .delete(`${process.env.DBAPI}/modules/${row.id}`, {
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
          sort:'-_updated',
          max_results: this.pagination.perPage,
          page: this.pagination.currentPage
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
        .get(`${process.env.DBAPI}/modules`, config)
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
<style scoped lang="scss">
.img-container {
  $size: 5rem;
  width: $size;
  height: $size;
  img {
    border-radius: 1rem;
  }
}
</style>