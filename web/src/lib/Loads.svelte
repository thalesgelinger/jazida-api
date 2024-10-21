<script lang="ts">
    import * as Table from "$lib/components/ui/table";
    import { onMount } from "svelte";
    import { api, ws } from "../service/api";
    import { Button } from "./components/ui/button";

    type Load = {
        id: string;
        client: string;
        plate: string;
        material: string;
        quantity: string;
        paymentMethod: string;
        signature: string;
    };

    type PaymentMethod = "CASH" | "INSTALLMENT";

    const getLoads = async (): Promise<Load[]> => {
        const response = await api.get("/loads", {
            headers: {
                Authorization: "admin",
            },
        });
        return response.data.map((d: { payment_method: PaymentMethod }) => {
            const load = {
                ...d,
                paymentMethod:
                    d.payment_method === "CASH" ? "A vista" : "A prazo",
            };

            // @ts-ignore
            delete load.payment_method;

            return load;
        }) as Load[];
    };

    let loads: Load[] = [];

    onMount(() => {
        getLoads().then((data) => {
            loads = data.reverse();
        });
    });

    ws.onmessage = ({ data }) => {
        const newLoad: any = JSON.parse(data);
        newLoad.paymentMethod =
            newLoad.payment_method === "CASH" ? "A vista" : "A prazo";
        newLoad.id = `${loads.length + 1}`;
        loads = [newLoad, ...loads];
    };

    function translateKey(key: keyof Load): string {
        switch (key) {
            case "id":
                return "ID";
            case "material":
                return "Material";
            case "plate":
                return "Placa";
            case "client":
                return "Cliente";
            case "quantity":
                return "Quantidade";
            case "signature":
                return "Assinatura";
            case "paymentMethod":
                return "Método de pagamento";
        }
    }

    function downloadCSV() {
        const header = [
            Object.keys(loads[0]).map((key) => translateKey(key as keyof Load)),
        ];
        const csv = header
            .concat(loads.map((item) => Object.values(item)))
            .map((row) => row.map((value) => `"${value}"`).join(","))
            .join("\n");
        const blob = new Blob([csv], { type: "text/csv" });
        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = "carregamentos.csv";

        link.click();
    }
</script>

<div class="rounded-md border">
    <div class="flex p-4 items-center">
        <h1 class="w-full text-center p-4 text-lg">Carregamentos</h1>
        <Button on:click={downloadCSV}>Baixar CSV</Button>
    </div>
    <Table.Root>
        <Table.Header>
            <Table.Row>
                <Table.Head>Id</Table.Head>
                <Table.Head>Cliente</Table.Head>
                <Table.Head>Placa</Table.Head>
                <Table.Head>Material</Table.Head>
                <Table.Head>Quantidade</Table.Head>
                <Table.Head>Método de pagamento</Table.Head>
                <Table.Head>Assinatura</Table.Head>
            </Table.Row>
        </Table.Header>
        <Table.Body>
            {#each loads as load}
                <Table.Row>
                    <Table.Cell>{load.id}</Table.Cell>
                    <Table.Cell>{load.client}</Table.Cell>
                    <Table.Cell>{load.plate}</Table.Cell>
                    <Table.Cell>{load.material}</Table.Cell>
                    <Table.Cell>{load.quantity}</Table.Cell>
                    <Table.Cell>{load.paymentMethod}</Table.Cell>
                    <Table.Cell>
                        <img
                            src={load.signature}
                            alt={`signature ${load.id}`}
                            class="h-10"
                        />
                    </Table.Cell>
                </Table.Row>
            {/each}
        </Table.Body>
    </Table.Root>
</div>
