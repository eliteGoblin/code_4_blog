FROM fluent/fluentd:v0.14.21-debian

ENV PATH /home/fluent/.gem/ruby/2.3.0/bin:$PATH

RUN apt-get -yqq update && \
 deps="sudo make gcc g++ libc-dev ruby-dev python-pip" && \
 apt-get install -yqq --no-install-recommends $deps && \
 gem install activesupport -v 5.1.4 && \
 gem install addressable -v 2.5.2 && \
 gem install bigdecimal -v 1.2.7 && \
 gem install concurrent-ruby -v 1.0.5 && \
 gem install cool.io -v 1.5.1 && \
 gem install did_you_mean -v 1.0.0 && \
 gem install domain_name -v 0.5.20170404 && \
 gem install fluent-plugin-kubernetes_metadata_filter -v 0.29.0 && \
 gem install gelf -v 3.0.0 && \
 gem install http -v 0.9.8 && \
 gem install http-cookie -v 1.0.3 && \
 gem install http-form_data -v 1.0.3 && \
 gem install http_parser.rb -v 0.6.0 && \
 gem install i18n -v 0.8.6 && \
 gem install io-console -v 0.4.5 && \
 gem install json -v 2.1.0 && \
 gem install json -v 1.8.3 && \
 gem install kubeclient -v 1.1.4 && \
 gem install lru_redux -v 1.1.0 && \
 gem install mime-types -v 3.1 && \
 gem install mime-types-data -v 3.2016.0521 && \
 gem install minitest -v 5.9.0 && \
 gem install msgpack -v 1.1.0 && \
 gem install net-telnet -v 0.1.1 && \
 gem install netrc -v 0.11.0 && \
 gem install oj -v 2.18.3 && \
 gem install power_assert -v 0.2.7 && \
 gem install psych -v 2.1.0 && \
 gem install public_suffix -v 3.0.0 && \
 gem install rake -v 10.5.0 && \
 gem install rdoc -v 4.2.1 && \
 gem install recursive-open-struct -v 1.0.0 && \
 gem install rest-client -v 2.0.2 && \
 gem install ruby_dig -v 0.0.2 && \
 gem install serverengine -v 2.0.5 && \
 gem install sigdump -v 0.2.4 && \
 gem install strptime -v 0.1.9 && \
 gem install test-unit -v 3.1.7 && \
 gem install thread_safe -v 0.3.6 && \
 gem install tzinfo -v 1.2.3 && \
 gem install tzinfo-data -v 1.2017.2 && \
 gem install unf -v 0.1.4 && \
 gem install unf_ext -v 0.0.7.4 && \
 gem install yajl-ruby -v 1.3.0 && \
 gem sources --clear-all && \
 SUDO_FORCE_REMOVE=yes apt-get purge -yqq --auto-remove -o APT::AutoRemove::RecommendsImportant=false $deps && \
 rm -rf /var/lib/apt/lists/* /home/fluent/.gem/ruby/2.3.0/cache/*.gem /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem

COPY ./data/fluent.conf /fluentd/etc
COPY plugins /fluentd/plugins/

RUN chmod -R 777 /var/log/

ENV LOG_PORT=53
ENV LISTENER_PORT=24224
ENV FLUENT_UID=0