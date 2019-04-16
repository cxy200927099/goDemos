package main

import (
	"context"
	"flag"
	"fmt"
	"goDemo/proto"
	"google.golang.org/grpc"
)

//const HOST = "127.0.0.1"

//const PORT = "20550"

//const TESTRECORD = 10

var (
	proxy = flag.String("proxy", "127.0.0.1:50051", "proxy")
)

var kRB5CONF string = ` 
[libdefaults]
default_realm = DATA.ONETHING.COM
dns_lookup_kdc = false
dns_lookup_realm = false
ticket_lifetime = 86400
renew_lifetime = 604800
forwardable = true
default_tgs_enctypes = aes256-cts
default_tkt_enctypes = aes256-cts
permitted_enctypes = aes256-cts
udp_preference_limit = 1
kdc_timeout = 3000
[realms]
DATA.ONETHING.COM = {
kdc = krb.data.onething.com
admin_server = krb.data.onething.com
default_domain = .data.onething.com
kdc = krb1.data.onething.com
}
[domain_realm]
.data.onething.com = DATA.ONETHING.COM
data.onething.com = DATA.ONETHING.COM`

func main() {

	//	startTime := currentTimeMillis()

	//logformatstr_ := "----%s\n"

	//logformatstr := "----%s 用时:%d-%d=%d毫秒\n\n"

	//logformattitle := "建立连接"

	//table := "video_finger_data"

	//rowkey := "bucket.0.group.0"

	//family := "cf"
	/*
	   	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	   	transport, err := thrift.NewTSocket(HOST + ":" + PORT)

	   	if err != nil {

	   		fmt.Println("failed when setup thrift socket! error=", err)
	   		return
	   	}

	   	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)

	   	if err := transport.Open(); err != nil {

	   		fmt.Println("failed when open thrift socket! error=", err)
	   		return
	   	}

	   	tmpendTime := currentTimeMillis()

	   	fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))

	   	defer transport.Close()

	   	//--------------Put

	   	logformattitle = "调用Put方法写数据"

	   	fmt.Printf(logformatstr_, logformattitle)

	   	startTime = currentTimeMillis()

	   	cvarr := []*hbase.TColumnValue{

	   		{

	   			Family: []byte(family),

	   			Qualifier: []byte("content"),

	   			Value: []byte("hello,world"),
	   		},
	   	}

	   	temptput := hbase.TPut{Row: []byte(rowkey), ColumnValues: cvarr}

	   	err = client.Put(context.Background(), []byte(table), &temptput)

	   	if err != nil {

	   		fmt.Printf("Put err:%s\n", err)

	   	} else {

	   		fmt.Println("Put done")

	   	}

	   	tmpendTime = currentTimeMillis()

	   	fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))

	   	//------------Get---------------

	   	logformattitle = "调用Get方法获取新增加的数据"

	   	fmt.Printf(logformatstr_, logformattitle)

	   	startTime = currentTimeMillis()

	   	result, err := client.Get(context.Background(), []byte(table), &hbase.TGet{Row: []byte(rowkey)})

	   	if err != nil {

	   		fmt.Printf("Get err:%s\n", err)

	   	} else {

	   		fmt.Println("Rowkey:" + string(result.Row))

	   		for _, cv := range result.ColumnValues {

	   			printscruct(cv)

	   		}

	   	}

	   	tmpendTime = currentTimeMillis()

	   	fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))
	   }

	   func currentTimeMillis() int64 {

	   	return time.Now().UnixNano() / 1000000

	   }

	   func printscruct(cv interface{}) {

	   	switch reflect.ValueOf(cv).Interface().(type) {

	   	case *hbase.TColumnValue:

	   		s := reflect.ValueOf(cv).Elem()

	   		typeOfT := s.Type()

	   		//获取Thrift2中struct的field

	   		for i := 0; i < s.NumField(); i++ {

	   			f := s.Field(i)

	   			fileldformatstr := "\t%d: %s(%s)= %v\n"

	   			switch f.Interface().(type) {

	   			case []uint8:

	   				fmt.Printf(fileldformatstr, i, typeOfT.Field(i).Name, f.Type(), string(f.Interface().([]uint8)))

	   			case *int64:

	   				var tempint64 int64

	   				if f.Interface().(*int64) == nil {

	   					tempint64 = 0

	   				} else {

	   					tempint64 = *f.Interface().(*int64)

	   				}

	   				fmt.Printf(fileldformatstr, i, typeOfT.Field(i).Name, f.Type(), tempint64)

	   			default:

	   				fmt.Print("I don't know")

	   			}

	   		}

	   	default:

	   		fmt.Print("I don't know")

	   		fmt.Print(reflect.ValueOf(cv))

	   	}
	*/
	/*
		spn := flag.String("spn", "", "spn name")
		username := flag.String("username", "hbase/tw06a1526", "username")
		realm := flag.String("realm", "DATA.ONETHING.COM", "relam")
		flag.Parse()
		fmt.Println("spn=", *spn, " username=", *username, " realm=", *realm)
		kt, err := keytab.Load("/home/root1/hbase.keytab")
		if err != nil {
			panic(err)
		}
		cl := client.NewClientWithKeytab(*username, *realm, kt)
		cfg, err := config.Load("/etc/krb5.conf")
		//fmt.Println(kRB5CONF)
		//cfg, err := config.NewConfigFromString(kRB5CONF)
		if err != nil {
			panic(err)
		}
		cl.WithConfig(cfg)
		if err := cl.Login(); err != nil {
			panic(err)
		}
		addr := "127.0.0.1:20550"
		rowkey := "bucket.0.group.0"
		cf := "cf:content"
		value := "hello"
		B64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		coder := base64.NewEncoding(B64Table)
		rowkeyBase64 := coder.EncodeToString([]byte(rowkey))
		cfBase64 := coder.EncodeToString([]byte(cf))
		valueBase64 := coder.EncodeToString([]byte(value))
		body := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`
		body += `<CellSet><Row key="` + rowkeyBase64 + `"><Cell column="` + cfBase64 + `">` + valueBase64 + `</Cell></Row></CellSet>`
		fmt.Println("request body:")
		fmt.Println(body)
		r, _ := http.NewRequest("PUT", "http://"+addr+"/video_finger_data/fakerow", strings.NewReader(body))
		r.Header.Add("Accept", "text/xml")
		r.Header.Add("Content-Type", "text/xml")
		r.Header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
		//	spn := "hbase/tw06a1526"
		//spn := "krbtgt/DATA.ONETHING.COM"
		if err := cl.SetSPNEGOHeader(r, *spn); err != nil {
			panic(err)
		}
		HTTPResp, err := http.DefaultClient.Do(r)
		/*
			t := &khttp.Transport{
				KeyTab:    "/home/root1/hbase.keytab",
				Principal: "hbase/tw06a1526@DATA.ONETHING.COM"}

			client := &http.Client{Transport: t}
			HTTPResp, err := client.Do(r)
	*/
	/*
		if err != nil {
			panic(err)
		}
		defer HTTPResp.Body.Close()
		bs, err := ioutil.ReadAll(HTTPResp.Body)
		if err != nil {
			panic(err)
		}
		println("status: ", HTTPResp.Status, "body:", string(bs))

		//get the cell
		getRequest, _ := http.NewRequest("GET", "http://"+addr+"/video_finger_data/"+rowkey+"/"+cf, nil)
		getRequest.Header.Add("Accept", "text/xml")
		if err := cl.SetSPNEGOHeader(getRequest, *spn); err != nil {
			panic(err)
		}
		HTTPGetResp, err := http.DefaultClient.Do(getRequest)
		if err != nil {
			panic(err)
		}
		defer HTTPGetResp.Body.Close()
		bs, err = ioutil.ReadAll(HTTPGetResp.Body)
		if err != nil {
			panic(err)
		}
		println("status: ", HTTPGetResp.Status, "body:", string(bs))
	*/
	flag.Parse()
	fmt.Println("proxy:", *proxy)
	conn, err := grpc.Dial(*proxy, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewHbaseServiceClient(conn)

    var (
		version uint32
		btype uint32
	)
	heartbeatRequest := proto.HbaseHeartBeatRequest{
		Version: &version, 
		Btype:   &btype,
	}
	heartBeatResp, err := client.HeartBeat(context.Background(), &heartbeatRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("heart beat resp: ", "version=", heartBeatResp.Version, " type=", heartBeatResp.Btype)
    
	keys := []string{"group.0.bucket.0","group.0.bucket.1","group.0.bucket.2","group.0.bucket.3","group.0.bucket.4"}
	datas := []string{"hello0", "hello1", "hello2", "hello3", "hello4"}
	dataArray := []*proto.Data{
		&proto.Data{
			Key:   &keys[0],
			Value: &datas[0],
		},
		&proto.Data{
			Key:   &keys[1],
			Value: &datas[1],
		},
		&proto.Data{
			Key:   &keys[2],
			Value: &datas[2],
		},
		&proto.Data{
			Key:   &keys[3],
			Value: &datas[3],
		},
		&proto.Data{
			Key:   &keys[4],
			Value: &datas[4],
		},
	}
	putRequest := proto.HbasePutRequest{
		Version:   &version,
		DataArray: dataArray,
	}
	putResp, err := client.Put(context.Background(), &putRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("put resp: ", "version=", putResp.Version, " status=", putResp.Status)

	getRequest := proto.HbaseGetRequest{
		Version: &version,
		Keys:    []string{"group.0.bucket.0", "group.0.bucket.1", "group.0.bucket.2", "group.0.bucket.3", "group.0.bucket.4"},
	}
	getResp, err := client.Get(context.Background(), &getRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("get resp: ", "version=", getResp.Version, " status=", getResp.Status, " data list:")
	for _, data := range getResp.DataArray {
		fmt.Println("key=", data.Key, " value=", data.Value)
	}
}
